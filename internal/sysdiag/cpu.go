package sysdiag

import (
	"fmt"
	"os/exec"
	"runtime"
)

// ShowCPULoad displays the current CPU load.
func ShowCPULoad() {
	if runtime.GOOS == "windows" {
		fmt.Println("CPU load metrics are not supported on Windows via this tool.")
	} else {
		output, err := exec.Command("sh", "-c", "uptime").Output()
		if err != nil {
			fmt.Printf("Failed to get CPU load: %s\n", err)
			return
		}
		fmt.Printf("Current CPU Load: %s\n", string(output))
	}
}
