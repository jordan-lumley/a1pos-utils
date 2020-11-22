package fs

import (
	"os"
	"path/filepath"
)

// EnsureFilePath ...
func EnsureFilePath(path string) error {
	dir, _ := filepath.Split(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}

	return nil
}
