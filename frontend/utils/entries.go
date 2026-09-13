package utils

import (
	"os"
	"slices"
	"time"
)

type Entry struct {
	Name    string
	IsDir   bool
	Size    int64
	ModTime time.Time
}

type Listing struct {
	Entries []Entry
	Folders int
	Files   int
	Size    int64
}

func ReadDir(dir string) (Listing, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return Listing{}, err
	}
	listing := Listing{Entries: make([]Entry, 0, len(dirEntries))}
	for _, dirEntry := range dirEntries {
		entry := Entry{Name: dirEntry.Name(), IsDir: dirEntry.IsDir()}
		if info, err := dirEntry.Info(); err == nil {
			entry.ModTime = info.ModTime()
			if !entry.IsDir {
				entry.Size = info.Size()
			}
		}
		if entry.IsDir {
			listing.Folders++
		} else {
			listing.Files++
			listing.Size += entry.Size
		}
		listing.Entries = append(listing.Entries, entry)
	}
	slices.SortStableFunc(listing.Entries, func(a, b Entry) int {
		switch {
		case a.IsDir == b.IsDir:
			return 0
		case a.IsDir:
			return -1
		default:
			return 1
		}
	})
	return listing, nil
}
