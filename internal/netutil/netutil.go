
package netutil

import "net"

func Contains(cidr string, ip net.IP) bool {
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	return network.Contains(ip)
}
