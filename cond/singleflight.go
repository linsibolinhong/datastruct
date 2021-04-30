package cond

import (
	"sync"
)

type call struct {
	value interface{}
	err error
}

type SingleFlight struct {
	mu sync.Mutex
	wg sync.WaitGroup
	calls map[string]*call
}

func (f *SingleFlight) doCall(key string, cal *call, fn func() (interface{}, error)) {
	cal.value, cal.err = fn()
	f.wg.Done()
	f.mu.Lock()
	delete(f.calls, key)
	f.mu.Unlock()
}

func (f *SingleFlight) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	f.mu.Lock()
	if f.calls == nil {
		f.calls = map[string]*call{}
	}

	callf, found := f.calls[key]
	if found {
		f.mu.Unlock()
		f.wg.Wait()
		return callf.value, callf.err
	}

	callf = new(call)
	f.wg.Add(1)
	f.calls[key] = callf
	f.mu.Unlock()
	f.doCall(key, callf, fn)
	return callf.value, callf.err
}
