package registry

import (
	"net"
	"path/filepath"
	"strings"

	"github.com/cloudwego/kitex/pkg/registry"
	registry_zookeeper "github.com/ishumei/krpc/registry-zookeeper"
)

type serviceInfo struct {
	ServiceName string
	Port        string
	*registry_zookeeper.NodeInfo
}

// path format as follows:
// /{ENV_ROLE}/{serviceName}/{ip}:{port}
func (n *serviceInfo) Path() string {
	var path string
	if !strings.HasPrefix(n.ServiceName, registry_zookeeper.Separator) {
		path = registry_zookeeper.Separator + n.ServiceName
	} else {
		path = n.ServiceName
	}
	return filepath.Join(path, net.JoinHostPort(n.Host, n.Port))
}

func newServiceInfo(info *registry.Info, inetPrefix string) *serviceInfo {
	if info == nil {
		return nil
	}
	if info.Addr == nil {
		return nil
	}
	host, port, err := net.SplitHostPort(info.Addr.String())
	if err != nil {
		return nil
	}
	if port == "" {
		return nil
	}
	ip := net.ParseIP(host)
	if ip == nil || ip.IsUnspecified() {
		host, err = GetLocalIp(inetPrefix)
		if err != nil {
			return nil
		}
	}
	return &serviceInfo{
		ServiceName: info.ServiceName,
		Port:        port,
		NodeInfo:    registry_zookeeper.NewNodeInfo(host, port),
	}
}
