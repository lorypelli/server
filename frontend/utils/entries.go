package utils

import (
	"io/fs"
	"os"
	"slices"
)

func Entries(dir string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	slices.SortStableFunc(entries, func(a, b fs.DirEntry) int {
		switch {
		case a.IsDir() == b.IsDir():
			return 0
		case a.IsDir():
			return -1
		default:
			return 1
		}
	})
	return entries, nil
}
