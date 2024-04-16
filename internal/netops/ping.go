package netops

import (
	"fmt"
	"net"
	"time"
)

func PingServer(host string) {
	ports := []string{"80", "443"}
	success := false

	for _, port := range ports {
		if tryConnect(host, port) {
			fmt.Printf("Ping to %s:%s successful.\n", host, port)
			success = true
			break
		}
	}

	if !success {
		fmt.Printf("Failed to ping %s: All tried ports refused connection or timed out.\n", host)
	}
}

func tryConnect(host, port string) bool {
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Printf("Failed to connect to %s on port %s: %s\n", host, port, err)
		return false
	}
	conn.Close()
	return true
}
