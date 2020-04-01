package event

import (
	"reflect"
	"sync"
)

type Emitter struct {
	listeners map[string]ListenerList
	mu        sync.RWMutex
}

type Listener func(...interface{})
type ListenerList []Listener

func (ls Listener) ptr() uintptr {
	return reflect.ValueOf(ls).Pointer()
}

func (l ListenerList) IndexOf(element Listener) int {
	p := element.ptr()
	for i := range l {
		if l[i].ptr() == p {
			return i
		}
	}
	return -1
}

func NewEmitter() *Emitter {
	return &Emitter{listeners: make(map[string]ListenerList)}
}

func (e *Emitter) emit(name string, arguments []interface{}) {
	e.mu.RLock()
	if listeners, ok := e.listeners[name]; ok {
		for _, l := range listeners {
			go l(arguments...)
		}
	}
	e.mu.RUnlock()
}

func (e *Emitter) Emit(name string, arguments ...interface{}) {
	e.emit(name, arguments)
}

func (e *Emitter) AddEventListener(name string, listener Listener) {
	e.mu.Lock()
	arr, ok := e.listeners[name]
	if !ok {
		t := make(ListenerList, 0, 10)
		e.listeners[name] = t
		arr = t
	}
	e.listeners[name] = append(arr, listener)
	e.mu.Unlock()
}

func (e *Emitter) On(name string, listener Listener) {
	e.AddEventListener(name, listener)
}

func (e *Emitter) RemoveEventListener(name string, listener Listener) {
	e.mu.Lock()
	if arr, ok := e.listeners[name]; ok {
		i := arr.IndexOf(listener)
		if i >= 0 {
			last := len(arr) - 1
			arr[i] = arr[last]
			arr[last] = nil
			e.listeners[name] = arr[:last]
		}
	}
	e.mu.Unlock()
}

func (e *Emitter) Off(name string, listener Listener) {
	e.RemoveEventListener(name, listener)
}
