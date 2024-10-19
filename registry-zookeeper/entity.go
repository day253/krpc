package registry_zookeeper

import (
	"os"
	"strconv"

	json "github.com/bytedance/sonic"
)

type serversetEndpoint struct {
	Host string
	Port int
}

type NodeInfo struct {
	// ishumei protocol
	Host string `json:"host"`
	Port int    `json:"port"`
	// kitex protocol
	Weight int               `json:"weight,omitempty"`
	Tags   map[string]string `json:"tags,omitempty"`
	// prometheus protocol
	// https://github.com/prometheus/prometheus/blob/b1566f761ef7d1c8ab6446afd017c0ce17e9224e/discovery/zookeeper/zookeeper.go#L250
	ServiceEndpoint     serversetEndpoint
	AdditionalEndpoints map[string]serversetEndpoint
	Status              string `json:"status"`
	Shard               int    `json:"shard"`
}

func (n *NodeInfo) Data() []byte {
	content, err := json.Marshal(n)
	if err != nil {
		return []byte{}
	}
	return content
}

func NewNodeInfo(host, port string) *NodeInfo {
	p, _ := strconv.Atoi(port)
	hostname, err := os.Hostname()
	if err != nil {
		hostname = host
	}
	return &NodeInfo{
		Host: host,
		Port: p, // rpc
		ServiceEndpoint: serversetEndpoint{
			Host: hostname,
			Port: p + 1, // http
		},
	}
}
