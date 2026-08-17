
package netutil

import "net"

// Contains reports whether ip falls within the CIDR range cidr.
// Containment is determined by applying the CIDR's network mask to
// both the network address and ip and comparing the results per
// RFC 4632 (i.e. bitwise, not byte/octet-wise).
func Contains(cidr string, ip net.IP) bool {
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	ip = ip.To4()
	if ip == nil {
		return false
	}
	return network.Contains(ip)
}
