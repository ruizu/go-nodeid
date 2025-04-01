package nodeid

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	netInterfaceAddrs = func() ([]net.Addr, error) {
		return nil, errors.New("test error")
	}

	id0, err := Get()
	assert.Error(t, err)
	assert.Equal(t, -1, id0)

	netInterfaceAddrs = func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{
				IP:   net.IPv4(127, 0, 0, 1),
				Mask: net.IPMask{255, 255, 255, 0},
			},
		}, nil
	}

	id1, err := Get()
	assert.Error(t, err)
	assert.Equal(t, -1, id1)

	netInterfaceAddrs = func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{
				IP:   net.IPv4(192, 168, 0, 120),
				Mask: net.IPMask{255, 255, 255, 0},
			},
		}, nil
	}
	id2, err := Get()
	assert.Equal(t, nil, err)
	assert.Equal(t, 120, id2)

	netInterfaceAddrs = func() ([]net.Addr, error) {
		return []net.Addr{
			&net.IPNet{
				IP:   net.IPv4(192, 168, 9, 120),
				Mask: net.IPMask{255, 255, 255, 0},
			},
		}, nil
	}
	id3, err := Get()
	assert.Equal(t, nil, err)
	assert.Equal(t, 2424, id3)
}
