package nodeid

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNoInterfaces(t *testing.T) {
	netInterfaceAddrs = func() ([]net.Addr, error) {
		return nil, nil
	}
	_, err := Get()
	assert.Error(t, err)
	assert.Equal(t, "no valid IPv4 address found", err.Error())
}

func TestGetLoopbackOnly(t *testing.T) {
	netInterfaceAddrs = func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{IP: net.ParseIP("127.0.0.1")},
		}, nil
	}
	_, err := Get()
	assert.Error(t, err)
	assert.Equal(t, "no valid IPv4 address found", err.Error())
}

func TestGetValidIPv4(t *testing.T) {
	netInterfaceAddrs = func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{IP: net.ParseIP("192.168.1.1")},
		}, nil
	}
	id, err := Get()
	assert.NoError(t, err)
	assert.Equal(t, int64(3232235777), id) // 192.168.1.1 -> 3232235777
}

func TestGetErrorFromNetInterfaceAddrs(t *testing.T) {
	netInterfaceAddrs = func() ([]net.Addr, error) {
		return nil, errors.New("mock error")
	}
	_, err := Get()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mock error")
}

func TestMustGetValidIPv4(t *testing.T) {
	netInterfaceAddrs = func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{IP: net.ParseIP("192.168.1.1")},
		}, nil
	}
	assert.NotPanics(t, func() {
		id := MustGet()
		assert.Equal(t, int64(3232235777), id) // 192.168.1.1 -> 3232235777
	})
}

func TestMustGetPanics(t *testing.T) {
	netInterfaceAddrs = func() ([]net.Addr, error) {
		return nil, nil
	}
	assert.Panics(t, func() {
		MustGet()
	})
}
