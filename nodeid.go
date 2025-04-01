package nodeid

import (
	"errors"
	"net"
)

var netInterfaceAddrs = net.InterfaceAddrs

func Get() (int, error) {
	addrs, err := netInterfaceAddrs()
	if err != nil {
		return -1, errors.New("unable to get addresses")
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}

		return (int(ip[2]) << 8) | int(ip[3]), nil
	}

	return -1, errors.New("unable to generate node id")
}
