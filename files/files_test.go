package files

import (
	"os"
	"testing"
)

// TestCreateTempDir tests CreateTempDir function
func TestCreateTempDir(t *testing.T) {
	// Create a temporary directory inside system temp directory
	dirName, err := CreateTempDir()

	// Check if the directory exists
	if !Exists(dirName) {
		t.Errorf("The directory %s does not exist", dirName)
	}

	// Write a file inside the directory
	fileName := dirName + "/test.txt"
	file, err := os.Create(fileName)
	if err != nil {
		t.Errorf("The file %s could not be created", fileName)
	}

	// Write some text inside the file
	_, err = file.WriteString("Hello World")
	if err != nil {
		t.Errorf("The file %s could not be written", fileName)
	}

	// Close the file
	err = file.Close()

	// Check if the file exists
	if !Exists(fileName) {
		t.Errorf("The file %s does not exist", fileName)
	}

	// Read the file content
	fileContent, err := ReadFile(fileName)
	if err != nil {
		t.Errorf("The file %s could not be read", fileName)
	}

	// Check the file content
	if fileContent != "Hello World" {
		t.Errorf("The file %s content is not correct", fileName)
	}

	// Delete the directory
	err = CleanupTempDirs()
	if err != nil {
		t.Errorf("The directory %s could not be deleted", dirName)
	}

	// Check if the directory exists
	if Exists(dirName) {
		t.Errorf("The directory %s exists", dirName)
	}
}
