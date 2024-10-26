package zookeeper

import (
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-zookeeper/zk"
)

type ZookeeperParams struct {
	servers        []string
	user           string
	password       string
	sessionTimeout time.Duration
	logger         zk.Logger
}

var (
	ErrorZkConnectedTimedOut = errors.New("timed out waiting for zk connected")
)

type Option func(c *ZookeeperParams)

func WithLogger(logger zk.Logger) Option {
	return func(c *ZookeeperParams) {
		c.logger = logger
	}
}

func (z *ZookeeperParams) Acl() []zk.ACL {
	if z.user != "" && z.password != "" {
		return zk.DigestACL(zk.PermAll, z.user, z.password)
	} else {
		return zk.WorldACL(zk.PermAll)
	}
}

func (z *ZookeeperParams) Connection() (*zk.Conn, error) {
	conn, event, err := zk.Connect(z.servers, z.sessionTimeout, zk.WithLogger(z.logger))
	if err != nil {
		return nil, err
	}
	if z.user != "" && z.password != "" {
		auth := []byte(fmt.Sprintf("%s:%s", z.user, z.password))
		if err := conn.AddAuth(Scheme, auth); err != nil {
			return nil, err
		}
	}
	ticker := time.NewTimer(time.Second * 10)
	for {
		select {
		case e := <-event:
			if e.State == zk.StateConnected {
				klog.Info("connected to zk server: ", z.servers)
				return conn, nil
			}
		case <-ticker.C:
			return nil, ErrorZkConnectedTimedOut
		}
	}
}

func NewZookeeperParams(servers []string, sessionTimeout time.Duration, user, password string) *ZookeeperParams {
	return &ZookeeperParams{
		servers:        servers,
		sessionTimeout: sessionTimeout,
		user:           user,
		password:       password,
		logger:         zk.DefaultLogger,
	}
}
