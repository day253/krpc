package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/go-zookeeper/zk"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
)

type ZookeeperResolver struct {
	*zk.Conn
}

// NewZookeeperResolverWithConf
func NewZookeeperResolverWithConf(c registry_zookeeper.Conf, options ...registry_zookeeper.Option) (*ZookeeperResolver, error) {
	return NewZookeeperResolverWithAuth(
		strings.Split(c.Metabase, registry_zookeeper.DefaultRegistrySeparater),
		time.Duration(c.TimeoutMs)*time.Millisecond,
		c.User,
		c.Password,
		options...,
	)
}

// NewZookeeperResolver create a zookeeper based resolver
func NewZookeeperResolver(servers []string, sessionTimeout time.Duration, options ...registry_zookeeper.Option) (*ZookeeperResolver, error) {
	return NewZookeeperResolverWithAuth(servers, sessionTimeout, "", "", options...)
}

// NewZookeeperResolver create a zookeeper based resolver with auth
func NewZookeeperResolverWithAuth(servers []string, sessionTimeout time.Duration, user, password string, options ...registry_zookeeper.Option) (*ZookeeperResolver, error) {
	p := registry_zookeeper.NewZookeeperParams(servers, sessionTimeout, user, password)
	for _, option := range options {
		option(p)
	}
	conn, err := p.Connection()
	if err != nil {
		return nil, err
	}
	return &ZookeeperResolver{
		Conn: conn,
	}, nil
}

// Target implements the Resolver interface.
func (r *ZookeeperResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) string {
	return target.ServiceName()
}

// Resolve implements the Resolver interface.
func (r *ZookeeperResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	path := desc
	if !strings.HasPrefix(path, registry_zookeeper.Separator) {
		path = registry_zookeeper.Separator + path
	}
	eps, err := r.getEndPoints(path)
	if err != nil {
		return discovery.Result{}, err
	}
	if len(eps) == 0 {
		return discovery.Result{}, fmt.Errorf("no instance remains for %v", desc)
	}
	instances, err := r.getInstances(eps, path)
	if err != nil {
		return discovery.Result{}, err
	}
	res := discovery.Result{
		Cacheable: true,
		CacheKey:  desc,
		Instances: instances,
	}
	return res, nil
}

func (r *ZookeeperResolver) getEndPoints(path string) ([]string, error) {
	child, _, err := r.Children(path)
	return child, err
}

func (r *ZookeeperResolver) detailEndPoints(path, ep string) (discovery.Instance, error) {
	data, _, err := r.Get(path + registry_zookeeper.Separator + ep)
	if err != nil {
		return nil, err
	}
	en := new(registry_zookeeper.NodeInfo)
	err = json.Unmarshal(data, en)
	if err != nil {
		return nil, fmt.Errorf("unmarshal data [%s] error, cause %w", data, err)
	}
	if en.Weight == 0 {
		en.Weight = 100
	}
	return discovery.NewInstance("tcp", ep, en.Weight, en.Tags), nil
}

func (r *ZookeeperResolver) getInstances(eps []string, path string) ([]discovery.Instance, error) {
	instances := make([]discovery.Instance, 0, len(eps))
	for _, ep := range eps {
		if host, port, err := net.SplitHostPort(ep); err == nil {
			if port == "" {
				return []discovery.Instance{}, fmt.Errorf("missing port when parse node [%s]", ep)
			}
			if host == "" {
				return []discovery.Instance{}, fmt.Errorf("missing host when parse node [%s]", ep)
			}
			ins, err := r.detailEndPoints(path, ep)
			if err != nil {
				return []discovery.Instance{}, fmt.Errorf("detail endpoint [%s] info error, cause %w", ep, err)
			}
			instances = append(instances, ins)
		} else {
			return []discovery.Instance{}, fmt.Errorf("parse node [%s] error, details info [%w]", ep, err)
		}
	}
	return instances, nil
}

// Diff implements the Resolver interface.
func (r *ZookeeperResolver) Diff(key string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(key, prev, next)
}

// Name implements the Resolver interface.
func (r *ZookeeperResolver) Name() string {
	return "zookeeper"
}
