package test

import (
	"fmt"
	"net"
)

const DEFAULT_LISTENER_ADDRESS = "0.0.0.0"

// GetFreePort asks the kernel for free open ports that are ready to use.
func GetFreePorts(count int) ([]int, error) {
	var ports []int
	for range count {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = l.Close()
		}()
		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
	}
	return ports, nil
}

func ListenerString(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}
