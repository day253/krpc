package zookeeper

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_selectLocalIp(t *testing.T) {
	testCases := []struct {
		name       string
		addresses  []*net.IPNet
		prefix     string
		expected   string
		shouldPass bool
	}{
		{
			name: "No matching prefix, first non-loopback IP",
			addresses: []*net.IPNet{
				{IP: net.ParseIP("192.168.1.100"), Mask: net.CIDRMask(24, 32)},
				{IP: net.ParseIP("10.0.2.15"), Mask: net.CIDRMask(8, 32)},
			},
			prefix:   "",
			expected: "192.168.1.100",
		},
		{
			name: "Matching prefix",
			addresses: []*net.IPNet{
				{IP: net.ParseIP("192.168.1.100"), Mask: net.CIDRMask(24, 32)},
				{IP: net.ParseIP("10.0.2.15"), Mask: net.CIDRMask(8, 32)},
			},
			prefix:   "10.0",
			expected: "10.0.2.15",
		},
		{
			name: "No IP with matching prefix",
			addresses: []*net.IPNet{
				{IP: net.ParseIP("192.168.1.100"), Mask: net.CIDRMask(24, 32)},
				{IP: net.ParseIP("10.0.1.15"), Mask: net.CIDRMask(8, 32)},
			},
			prefix:   "10.1",
			expected: "",
		},
		{
			name: "All loopback addresses",
			addresses: []*net.IPNet{
				{IP: net.ParseIP("127.0.0.1"), Mask: net.CIDRMask(8, 32)},
				{IP: net.ParseIP("127.0.0.2"), Mask: net.CIDRMask(8, 32)},
			},
			prefix:   "10.0",
			expected: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := selectLocalIp(func() []net.Addr {
				addrs := make([]net.Addr, len(tc.addresses))
				for i, ipnet := range tc.addresses {
					addrs[i] = ipnet
				}
				return addrs
			}(), tc.prefix)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
