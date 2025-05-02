package nodeid

import (
	"fmt"
	"net"
)

// Add a mockable variable for net.InterfaceAddrs
var netInterfaceAddrs func() ([]net.Addr, error)

func init() {
	netInterfaceAddrs = net.InterfaceAddrs
}

// Get retrieves a unique identifier based on the first non-loopback IPv4 address of the machine.
// It returns the identifier as an int64 and an error if no valid IPv4 address is found.
func Get() (int64, error) {
	addrs, err := netInterfaceAddrs()
	if err != nil {
		return 0, fmt.Errorf("failed to get network interfaces: %w", err)
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip := ipNet.IP.To4()
				id := int64(ip[0])<<24 | int64(ip[1])<<16 | int64(ip[2])<<8 | int64(ip[3])
				return id, nil
			}
		}
	}

	return 0, fmt.Errorf("no valid IPv4 address found")
}

// MustGet retrieves a unique identifier based on the first non-loopback IPv4 address of the machine.
// It panics if no valid IPv4 address is found or if an error occurs.
func MustGet() int64 {
	id, err := Get()
	if err != nil {
		panic(err)
	}
	return id
}
