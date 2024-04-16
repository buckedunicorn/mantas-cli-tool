package fileops

import (
	"fmt"
)

// HandleFileOperations routes file operation commands to the appropriate functions.
func HandleFileOperations(args []string) {
	if len(args) < 2 {
		fmt.Println("Not enough arguments for file operations")
		return
	}

	command := args[0]
	filename := args[1]
	content := ""

	if len(args) > 2 {
		content = args[2]
	}

	switch command {
	case "create":
		CreateFile(filename)
	case "read":
		ReadFile(filename)
	case "update":
		UpdateFile(filename, content)
	case "delete":
		DeleteFile(filename)
	default:
		fmt.Printf("Unknown file operation: %s\n", command)
	}
}
