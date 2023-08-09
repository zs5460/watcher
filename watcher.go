// Package watcher implements a simple file monitor that calls a callback
// function when a file is modified, typically used to automatically load
// configuration files
package watcher

import (
	"os"
	"time"
)

type watcher struct {
	filename string
	onChange func()
}

func New(fn string, call func()) (*watcher, error) {
	_, err := os.Stat(fn)
	if err != nil {
		return nil, err
	}
	return &watcher{filename: fn, onChange: call}, nil
}

func Must(fn string, call func()) *watcher {
	w, err := New(fn, call)
	if err != nil {
		panic(err)
	}
	return w
}

func (wc *watcher) watch() {
	info, _ := os.Stat(wc.filename)
	lastModTime := info.ModTime()
	for range time.Tick(1 * time.Second) {
		info, _ := os.Stat(wc.filename)
		if !info.ModTime().After(lastModTime) {
			continue
		}
		lastModTime = info.ModTime()
		wc.onChange()
	}
}

func (wc *watcher) Start() {
	go wc.watch()
}
