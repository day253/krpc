package zookeeper

import (
	"net"
	"path/filepath"
	"strings"

	"github.com/cloudwego/kitex/pkg/registry"
)

type serviceInfo struct {
	ServiceName string
	Port        string
	*NodeInfo
}

// path format as follows:
// /{ENV_ROLE}/{serviceName}/{ip}:{port}
func (n *serviceInfo) Path() string {
	var path string
	if !strings.HasPrefix(n.ServiceName, Separator) {
		path = Separator + n.ServiceName
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
		NodeInfo:    NewNodeInfo(host, port),
	}
}
