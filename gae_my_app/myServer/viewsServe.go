package myServer

import (
	"os"
	"path/filepath"
)

func GetFileName(path string) ([]string, error) {
	var names []string
	err := filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
		if file == nil {
			return err
		}
		if !file.IsDir() {
			names = append(names, "./views/"+file.Name())
		}
		return err

	})
	return names, err
}
