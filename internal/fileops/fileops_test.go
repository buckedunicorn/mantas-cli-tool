package fileops

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	filename := "testfile.txt"
	CreateFile(filename)
	defer os.Remove(filename)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("CreateFile failed to create file")
	}
}

func TestReadFile(t *testing.T) {
	filename := "testfile.txt"
	content := "Hello, world!"
	ioutil.WriteFile(filename, []byte(content), 0644)
	defer os.Remove(filename)

	ReadFile(filename)
	// Output checks would require a more complex setup with capturing stdout
}

func TestUpdateFile(t *testing.T) {
	filename := "testfile.txt"
	// initialContent := "Hello"
	updatedContent := "Hello, world!"
	CreateFile(filename)
	defer os.Remove(filename)

	UpdateFile(filename, updatedContent)
	data, _ := ioutil.ReadFile(filename)
	if string(data) != updatedContent {
		t.Errorf("UpdateFile did not correctly update the file")
	}
}

func TestDeleteFile(t *testing.T) {
	filename := "testfile.txt"
	CreateFile(filename)
	DeleteFile(filename)

	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		t.Errorf("DeleteFile failed to delete the file")
	}
}
