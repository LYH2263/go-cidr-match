
package netutil

import (
	"net"
	"testing"
)

func TestSlash24(t *testing.T) {
	if !Contains("192.168.1.0/24", net.ParseIP("192.168.1.50")) {
		t.Fatal("should contain")
	}
	if Contains("192.168.1.0/24", net.ParseIP("192.168.2.50")) {
		t.Fatal("should not contain")
	}
}
