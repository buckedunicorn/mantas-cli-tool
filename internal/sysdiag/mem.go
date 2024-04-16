package sysdiag

import (
	"fmt"
	"runtime"
)

// ShowMemoryUsage displays the current memory usage statistics.
func ShowMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Allocated memory: %v bytes\n", m.Alloc)
	fmt.Printf("Total memory (Sys): %v bytes\n", m.Sys)
	fmt.Printf("Number of mallocs: %v\n", m.Mallocs)
	fmt.Printf("Number of frees: %v\n", m.Frees)
}
