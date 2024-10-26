package example

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/ishumei/krpc/zookeeper"
	"github.com/stretchr/testify/assert"
)

var (
	testZkServers    = []string{"127.0.0.1:2181"}
	testRegisterPath = "/public/test"
)

func testRegisterPort() int {
	rand.Seed(time.Now().UnixNano())
	minPort := 1
	maxPort := 65535
	return rand.Intn(maxPort-minPort+1) + minPort
}

func TestZookeeperDiscovery(t *testing.T) {
	// port
	targetPort := testRegisterPort()

	// register
	r, err := zookeeper.NewZookeeperRegistry(testZkServers, 40*time.Second, "")
	assert.Nil(t, err)
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", targetPort))
	assert.Nil(t, err)
	info := &registry.Info{ServiceName: testRegisterPath, Weight: 100, PayloadCodec: "thrift", Addr: addr}
	assert.Nil(t, r.Register(info))

	// resolve
	res, err := zookeeper.NewZookeeperResolver(testZkServers, 40*time.Second)
	assert.Nil(t, err)
	target := res.Target(context.Background(), rpcinfo.NewEndpointInfo(testRegisterPath, "", nil, nil))
	result, err := res.Resolve(context.Background(), target)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(result.Instances))

	// compare data
	instance := result.Instances[0]
	host, port, err := net.SplitHostPort(instance.Address().String())
	assert.Nil(t, err)
	assert.NotEmpty(t, host)
	assert.Equal(t, fmt.Sprintf("%d", targetPort), port)
	assert.Equal(t, instance.Weight(), info.Weight)

	// deregister
	assert.Nil(t, r.Deregister(info))

	// resolve again
	result, err = res.Resolve(context.Background(), target)
	assert.EqualError(t, err, fmt.Sprintf("no instance remains for %s", testRegisterPath))
}

func TestZookeeperResolverWithAuth(t *testing.T) {
	// port
	targetPort := testRegisterPort()

	// register
	r, err := zookeeper.NewZookeeperRegistryWithAuth(testZkServers, 40*time.Second, "zkadmin", "zkadmin123", "")
	assert.Nil(t, err)
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", targetPort))
	assert.Nil(t, err)
	info := &registry.Info{ServiceName: testRegisterPath, Weight: 100, PayloadCodec: "thrift", Addr: addr}
	assert.Nil(t, r.Register(info))

	// resolve
	res, err := zookeeper.NewZookeeperResolverWithAuth(testZkServers, 40*time.Second, "zkadmin", "zkadmin123")
	assert.Nil(t, err)
	target := res.Target(context.Background(), rpcinfo.NewEndpointInfo(testRegisterPath, "", nil, nil))
	result, err := res.Resolve(context.Background(), target)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(result.Instances))

	// compare data
	instance := result.Instances[0]
	host, port, err := net.SplitHostPort(instance.Address().String())
	assert.Nil(t, err)
	assert.NotEmpty(t, host)
	assert.Equal(t, fmt.Sprintf("%d", targetPort), port)
	assert.Equal(t, instance.Weight(), info.Weight)

	// deregister
	assert.Nil(t, r.Deregister(info))

	// resolve again
	result, err = res.Resolve(context.Background(), target)
	assert.EqualError(t, err, fmt.Sprintf("no instance remains for %s", testRegisterPath))
}
