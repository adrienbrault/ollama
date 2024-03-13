package server

import (
	"os"
	"path/filepath"
	"strings"
)

// fixOldBlobNames renames all blobs with (":") in the name to use ("-")
// instead.
func fixOldBlobNames() error {
	path, err := GetBlobsPath("")
	if err != nil {
		return err
	}
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, ":") {
			newPath := strings.ReplaceAll(path, ":", "-")
			return os.Rename(path, newPath)
		}
		return nil
	})
}
