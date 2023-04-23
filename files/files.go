package files

import (
	"fmt"
	"github.com/fkmatsuda-dev/commons/random"
	"io"
	"os"
)

var (
	tempDirs []string
)

// CreateTempDir creates a temporary directory inside the system temp folder
func CreateTempDir() (string, error) {
	// get system temp folder
	tempFolder := os.TempDir()

	// create a random directory name
	dirName := random.String(10)

	// format the directory name
	dirName = fmt.Sprintf("%s%s%s", tempFolder, string(os.PathSeparator), dirName)

	// create a temporary directory inside system temp directory
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		return "", err
	}
	// add the temporary directory to the list of temporary directories
	tempDirs = append(tempDirs, dirName)
	// return the temporary directory
	return dirName, nil

}

// DeleteTempDir deletes a temporary directory
func DeleteTempDir(dir string) error {
	// check if the directory is a temporary directory

	// delete the temporary directory
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	// remove the temporary directory from the list of temporary directories
	for i, tempDir := range tempDirs {
		if tempDir == dir {
			tempDirs = append(tempDirs[:i], tempDirs[i+1:]...)
			break
		}
	}
	return nil
}

// CleanupTempDirs deletes all temporary directories
func CleanupTempDirs() error {
	// iterate over the temporary directories
	for _, tempDir := range tempDirs {
		// delete the temporary directory
		err := DeleteTempDir(tempDir)
		if err != nil {
			return err
		}
	}
	return nil
}

// Exists checks if a file or directory exists
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// ReadFile reads a file and returns its content
func ReadFile(fileName string) (string, error) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	// Close the file
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Return the file content
	return string(fileContent), nil
}
