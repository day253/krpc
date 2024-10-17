package registry

import (
	"errors"
	"net"
	"strings"
)

var (
	ErrorNotFoundIPv4Address = errors.New("not found ipv4 address")
)

func GetLocalIp(prefix string) (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", ErrorNotFoundIPv4Address
	}
	return selectLocalIp(addrs, prefix)
}

func selectLocalIp(addresses []net.Addr, prefix string) (string, error) {
	for _, address := range addresses {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if len(prefix) > 0 {
					if strings.HasPrefix(ipnet.IP.String(), prefix) {
						return ipnet.IP.String(), nil
					}
				} else {
					return ipnet.IP.String(), nil
				}
			}
		}
	}
	return "", ErrorNotFoundIPv4Address
}
