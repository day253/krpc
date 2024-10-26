package zookeeper

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/go-zookeeper/zk"
	"go.uber.org/multierr"
)

const (
	maxDeleteRetries = 5
)

type meta struct {
	path     string
	data     []byte
	wg       sync.WaitGroup
	canceler context.CancelFunc
}

type ZookeeperRegistry struct {
	*ZookeeperParams
	sync.RWMutex
	inetPrefix string
	metas      map[string]*meta
	conn       *zk.Conn
	connMutex  sync.RWMutex
	interval   time.Duration
}

var (
	ErrorConnection       = errors.New("error connection")
	ErrorServiceInfo      = errors.New("empty service info")
	ErrorMaxDeleteRetries = errors.New("max retries reached for node delete")
)

// NewZookeeperRegistryWithConf
func NewZookeeperRegistryWithConf(c Conf, inetPrefix string, options ...Option) (*ZookeeperRegistry, error) {
	return NewZookeeperRegistryWithAuth(
		strings.Split(c.Metabase, DefaultRegistrySeparater),
		time.Duration(c.TimeoutMs)*time.Millisecond,
		c.User,
		c.Password,
		inetPrefix,
		options...,
	)
}

func NewZookeeperRegistry(servers []string, sessionTimeout time.Duration, inetPrefix string, options ...Option) (*ZookeeperRegistry, error) {
	return NewZookeeperRegistryWithAuth(servers, sessionTimeout, "", "", inetPrefix, options...)
}

func NewZookeeperRegistryWithAuth(servers []string, sessionTimeout time.Duration, user, password string, inetPrefix string, options ...Option) (*ZookeeperRegistry, error) {
	p := NewZookeeperParams(servers, sessionTimeout, user, password)
	for _, option := range options {
		option(p)
	}
	conn, err := p.Connection()
	if err != nil {
		return nil, err
	}
	return &ZookeeperRegistry{
		ZookeeperParams: p,
		conn:            conn,
		inetPrefix:      inetPrefix,
		metas:           make(map[string]*meta),
		interval:        maxTimeDuration(sessionTimeout, time.Second*30),
	}, nil
}

func maxTimeDuration(a, b time.Duration) time.Duration {
	if a > b {
		return a
	}
	return b
}

func (z *ZookeeperRegistry) Register(info *registry.Info) error {
	z.Lock()
	defer z.Unlock()
	if info == nil {
		// 为了方便实现zk-online的功能，将保存的所有节点都重新上线
		return z.online()
	}
	e := newServiceInfo(info, z.inetPrefix)
	if e == nil {
		return ErrorServiceInfo
	}
	return z.registerNode(e.Path(), e.Data())
}

func (z *ZookeeperRegistry) Deregister(info *registry.Info) error {
	z.Lock()
	defer z.Unlock()
	if info == nil {
		// 为了方便实现zk-offline的功能，将保存的所有节点都下线
		return z.offline()
	}
	e := newServiceInfo(info, z.inetPrefix)
	if e == nil {
		return ErrorServiceInfo
	}
	return z.deregisterNode(e.Path(), e.Data())
}

func (z *ZookeeperRegistry) Check() error {
	z.RLock()
	defer z.RUnlock()
	var err error
	for _, m := range z.metas {
		if exists, _, _ := z.exists(m.path); !exists {
			err = multierr.Append(err, fmt.Errorf("%s not exists", m.path))
		}
	}
	return err
}

func (z *ZookeeperRegistry) online() error {
	var errs error
	for _, m := range z.metas {
		if err := z.registerNode(m.path, m.data); err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}

func (z *ZookeeperRegistry) offline() error {
	var errs error
	for _, m := range z.metas {
		if err := z.deregisterNode(m.path, m.data); err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}

func (z *ZookeeperRegistry) registerNode(path string, data []byte) error {
	klog.Info("register path: ", path, ", data: ", string(data))
	if err := z.createNode(path, data, zk.FlagEphemeral); err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	m := &meta{
		path:     path,
		data:     data,
		canceler: cancel,
	}
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		z.keepalive(ctx, path, data)
	}()
	z.metas[path] = m
	return nil
}

func (z *ZookeeperRegistry) deregisterNode(path string, data []byte) error {
	klog.Info("deregister path: ", path, ", data: ", string(data))
	if m, ok := z.metas[path]; ok {
		m.canceler()
		m.wg.Wait()
	}
	if err := z.deleteNodeWithRetry(path, -1); err != nil && err != zk.ErrNoNode {
		return err
	}
	return nil
}

// keepalive re-register data node info when bad connection recovered
func (z *ZookeeperRegistry) keepalive(ctx context.Context, path string, data []byte) {
	klog.Info("keepalive path: ", path)
	var err error
retry:
	if err != nil {
		time.Sleep(10 * time.Second)
		if err = z.reConnect(); err != nil {
			klog.Info("keepalive reConnect err: ", err)
			goto retry
		}
		if err = z.createNode(path, data, zk.FlagEphemeral); err != nil {
			klog.Info("keepalive createNode err: ", err)
			goto retry
		}
	}
	ticker := time.NewTicker(z.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			err := z.deleteNodeWithRetry(path, -1)
			klog.Info("keepalive loop deleteNode err: ", err)
			return
		case <-ticker.C:
			var exists bool
			exists, _, err = z.exists(path)
			if err != nil {
				klog.Error("exists node error: ", err)
				goto retry
			}
			if !exists {
				klog.Info("re-ensure path: ", path, ", data:", string(data))
				err = z.createNode(path, data, zk.FlagEphemeral)
				if err != nil {
					klog.Error("create node error: ", err)
					goto retry
				}
			}
		}
	}
}

// createNode ensure node exists, if not exist, create and set data
func (z *ZookeeperRegistry) createNode(path string, data []byte, flags int32) error {
	exists, stat, err := z.exists(path)
	if err != nil {
		return err
	}
	// ephemeral nodes handling after restart
	// fixes a race condition if the server crashes without using CreateProtectedEphemeralSequential()
	// https://github.com/go-kratos/kratos/blob/main/contrib/registry/zookeeper/register.go
	if exists && flags&zk.FlagEphemeral == zk.FlagEphemeral {
		err = z.delete(path, stat.Version)
		if err != nil && err != zk.ErrNoNode {
			return err
		}
		exists = false
	}
	if !exists {
		_, err = z.create(path, data, flags, z.Acl())
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *ZookeeperRegistry) deleteNodeWithRetry(path string, version int32) error {
	for i := 0; i < maxDeleteRetries; i++ {
		err := z.delete(path, version)
		if err == nil || err == zk.ErrNoNode {
			return err
		}
		// Exponential backoff
		if i < maxDeleteRetries-1 {
			time.Sleep(time.Second * time.Duration(math.Pow(2, float64(i))))
		}
	}
	return ErrorMaxDeleteRetries
}

func (z *ZookeeperRegistry) reConnect() error {
	z.connMutex.Lock()
	defer z.connMutex.Unlock()
	if z.conn != nil {
		z.conn.Close()
		z.conn = nil
	}
	conn, err := z.Connection()
	if err != nil {
		return err
	}
	z.conn = conn
	return nil
}

func (z *ZookeeperRegistry) exists(path string) (bool, *zk.Stat, error) {
	z.connMutex.RLock()
	defer z.connMutex.RUnlock()
	if z.conn == nil {
		return false, nil, ErrorConnection
	}
	return z.conn.Exists(path)
}

func (z *ZookeeperRegistry) create(path string, data []byte, flags int32, acl []zk.ACL) (string, error) {
	z.connMutex.RLock()
	defer z.connMutex.RUnlock()
	if z.conn == nil {
		return "", ErrorConnection
	}
	return z.conn.Create(path, data, flags, acl)
}

func (z *ZookeeperRegistry) delete(path string, version int32) error {
	z.connMutex.RLock()
	defer z.connMutex.RUnlock()
	if z.conn == nil {
		return ErrorConnection
	}
	return z.conn.Delete(path, version)
}

func (z *ZookeeperRegistry) Children(path string) ([]string, *zk.Stat, error) {
	z.connMutex.RLock()
	defer z.connMutex.RUnlock()
	if z.conn == nil {
		return []string{}, nil, ErrorConnection
	}
	return z.conn.Children(path)
}
