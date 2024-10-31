package internal

import (
	"net"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if address, ok := addr.(*net.IPNet); ok && address.IP.To4() != nil && address.IP.IsPrivate() {
				return address.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
