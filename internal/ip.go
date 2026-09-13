package internal

import "net"

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return LocalIP
	}
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && ip.IP.To4() != nil && ip.IP.IsPrivate() {
			return ip.IP.String()
		}
	}
	return LocalIP
}
