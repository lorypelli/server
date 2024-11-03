package templ

import "os"

type File struct {
	name  string
	isDir bool
}

func GetFiles(dir string) []File {
	var f []File
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		f = append(f, File{
			name:  file.Name(),
			isDir: file.IsDir(),
		})
	}
	return f
}
