package utils

import (
	"path/filepath"

	"github.com/a-h/templ"
)

func GetURL(path string, dir string, f File) templ.SafeURL {
	if path == "/" {
		return templ.URL(f.Name)
	}
	return templ.URL(filepath.Base(dir) + "/" + f.Name)
}
