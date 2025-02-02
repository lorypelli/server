package pkg

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Monitor(path string, hasChanged chan bool) {
	defer close(hasChanged)
	times := make(map[string]time.Time)
	filepath.WalkDir(path, func(path string, _ fs.DirEntry, _ error) error {
		if strings.HasPrefix(path, ".") {
			return nil
		}
		fileInfo, _ := os.Stat(path)
		times[path] = fileInfo.ModTime()
		return nil
	})
	for {
		time.Sleep(time.Second)
		filepath.WalkDir(path, func(path string, _ fs.DirEntry, _ error) error {
			if strings.HasPrefix(path, ".") {
				return nil
			}
			fileInfo, _ := os.Stat(path)
			newTime := fileInfo.ModTime()
			if !newTime.Equal(times[path]) {
				times[path] = newTime
				hasChanged <- true
			}
			for path := range times {
				if _, err := os.Stat(path); err != nil {
					delete(times, path)
					hasChanged <- true
				}
			}
			return nil
		})
	}
}
