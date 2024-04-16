package sysdiag

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	return out
}

func TestShowMemoryUsage(t *testing.T) {
	output := captureOutput(ShowMemoryUsage)
	if output == "" {
		t.Error("Expected memory usage output, got none")
	} else {
		fmt.Println("Memory Usage Output:", output)
	}
}

func TestShowCPULoad(t *testing.T) {
	output := captureOutput(ShowCPULoad)
	if output == "" {
		t.Error("Expected CPU load output, got none")
	} else {
		fmt.Println("CPU Load Output:", output)
	}
}
