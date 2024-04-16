package sysdiag

import (
	"fmt"
)

func HandleSystemDiagnostics(args []string) {
	if len(args) < 1 {
		fmt.Println("Not enough arguments for system diagnostics")
		return
	}

	command := args[0]

	switch command {
	case "mem":
		ShowMemoryUsage()
	case "cpu":
		ShowCPULoad()
	default:
		fmt.Printf("Unknown system diagnostic operation: %s\n", command)
	}
}
