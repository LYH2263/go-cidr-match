
package netutil

import "net"

func Contains(cidr string, ip net.IP) bool {
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	ip = ip.To4()
	if ip == nil {
		return false
	}
	// BUG: only compare first octet
	return ip[0] == network.IP.To4()[0]
}
