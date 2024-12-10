package goutil

import (
	"os"
	"path/filepath"
)

func File_WalkDir(dir string) map[string]string {
	files := make(map[string]string)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files[path] = filepath.Base(path)
		}
		return nil
	})
	return files
}
