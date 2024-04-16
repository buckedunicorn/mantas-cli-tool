package fileops

import (
	"fmt"
	"os"
)

// DeleteFile removes the specified file.
func DeleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Printf("Failed to delete file: %s\n", err)
		return
	}
	fmt.Println("File deleted successfully:", filename)
}
