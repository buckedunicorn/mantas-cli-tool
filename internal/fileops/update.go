package fileops

import (
	"fmt"
	"os"
)

func UpdateFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Failed to open file for updating: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Failed to write to file: %s\n", err)
		return
	}
	fmt.Println("File updated successfully:", filename)
}
