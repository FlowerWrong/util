package util

import (
	"sync"
)

// WaitGroupWrapper struct
type WaitGroupWrapper struct {
	sync.WaitGroup
}

// Wrap a callback
func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
