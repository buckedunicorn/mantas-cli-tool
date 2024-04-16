package fileops

import (
	"fmt"
	"io/ioutil"
)

// ReadFile reads and displays the contents of the specified file.
func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", err)
		return
	}
	fmt.Println("File contents:", string(data))
}
