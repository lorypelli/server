package pkg

import (
	"io/fs"
	"maps"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type watcher struct {
	dir  string
	mu   sync.Mutex
	subs map[chan struct{}]struct{}
}

func watch(dir string) *watcher {
	w := &watcher{dir: dir, subs: make(map[chan struct{}]struct{})}
	go w.run()
	return w
}

func (w *watcher) subscribe() (<-chan struct{}, func()) {
	ch := make(chan struct{}, 1)
	w.mu.Lock()
	w.subs[ch] = struct{}{}
	w.mu.Unlock()
	return ch, func() {
		w.mu.Lock()
		delete(w.subs, ch)
		w.mu.Unlock()
	}
}

func (w *watcher) run() {
	var previous map[string]time.Time
	for range time.Tick(time.Second) {
		if !w.subscribed() {
			previous = nil
			continue
		}
		current := snapshot(w.dir)
		if previous != nil && !maps.EqualFunc(previous, current, time.Time.Equal) {
			w.broadcast()
		}
		previous = current
	}
}

func (w *watcher) subscribed() bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	return len(w.subs) > 0
}

func (w *watcher) broadcast() {
	w.mu.Lock()
	defer w.mu.Unlock()
	for ch := range w.subs {
		select {
		case ch <- struct{}{}:
		default:
		}
	}
}

func snapshot(dir string) map[string]time.Time {
	times := make(map[string]time.Time)
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if path != dir && strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if info, err := d.Info(); err == nil {
			times[path] = info.ModTime()
		}
		return nil
	})
	return times
}
