package utils

import (
	"os"
)

type File struct {
	Name  string
	IsDir bool
}

func GetFiles(dir string) []File {
	var f []File
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		f = append(f, File{
			Name:  file.Name(),
			IsDir: file.IsDir(),
		})
	}
	return f
}
