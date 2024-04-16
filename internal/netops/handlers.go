package netops

import (
	"fmt"
)

func HandleNetworkOperations(args []string) {
	if len(args) < 1 {
		fmt.Println("Not enough arguments for network operations")
		return
	}

	command := args[0]
	target := ""

	if len(args) > 1 {
		target = args[1]
	}

	switch command {
	case "ping":
		if target == "" {
			fmt.Println("Ping command requires a host argument")
		} else {
			PingServer(target)
		}
	case "fetch":
		if target == "" {
			fmt.Println("Fetch command requires a URL argument")
		} else {
			FetchURL(target)
		}
	default:
		fmt.Printf("Unknown network operation: %s\n", command)
	}
}
