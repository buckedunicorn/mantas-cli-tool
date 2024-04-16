package fileops

import (
	"fmt"
	"os"
)

// CreateFile creates a new file with the specified name.
func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file: %s\n", err)
		return
	}
	defer file.Close()
	fmt.Println("File created successfully:", filename)
}
