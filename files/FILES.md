# "FILES" Package

## Installation
To install the "files" package, run the following command:
```bash
go get github.com/fkmatsuda-dev/commons/files
```
## Usage
To use the "files" package, import it in your Go code:
```go
import "github.com/fkmatsuda-dev/commons/files"
```

## Temporary Directories
The "files" package provides a function to create a temporary directory and a function to delete a directory and all its contents. The temporary directory is created in the system's temporary directory and its name is a random string of 32 characters. The temporary directory is deleted when call the function "files.CleanupTempDirs()"

Tos create a temporary directory, use the files.CreateTempDir() function:
```go
tempDir, err := files.CreateTempDir()
if err != nil {
    // Handle error
}
```

To delete a directory and all its contents, use the files.DeleteTempDir() function:
```go
err := files.DeleteTempDir(tempDir)
if err != nil {
    // Handle error
}
```

To delete all temporary directories created by the "files" package, use the files.CleanupTempDirs() function:
```go
err := files.CleanupTempDirs()
if err != nil {
    // Handle error
}
```
