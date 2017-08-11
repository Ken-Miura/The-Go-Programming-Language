// Copyright 2017 Ken Miura
package memo

import (
	"errors"
	"sync"
)

//!+Func

// Func is the type of the function to memoize.
type Func func(key string, cancel <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
	complete chan struct{}
	cancel   <-chan struct{}
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancel <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, make(chan struct{}), cancel}
	select {
	case res := <-response:
		return res.value, res.err
	case <-cancel:
		return nil, errors.New("operation canceled")
	}
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

type cache struct {
	sync.Mutex
	mapping map[string]*entry
}

func (c *cache) get(key string) (*entry, bool) {
	defer c.Unlock()
	c.Lock()
	e, ok := c.mapping[key]
	return e, ok
}

func (c *cache) put(key string, e *entry) {
	defer c.Unlock()
	c.Lock()
	c.mapping[key] = e
}

func (c *cache) delete(key string) {
	defer c.Unlock()
	c.Lock()
	delete(c.mapping, key)
}

//!+monitor
func (memo *Memo) server(f Func) {
	cacheMap := &cache{sync.Mutex{}, make(map[string]*entry)}
	for req := range memo.requests {
		e, _ := cacheMap.get(req.key)
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cacheMap.put(req.key, e)
			go e.call(f, req.key, req.cancel) // call f(key)
		}
		go func() {
			select {
			case <-req.cancel:
				cacheMap.delete(req.key)
			case <-req.complete:
				// do nothing
			}
		}()
		go func() {
			e.deliver(req.response)
			close(req.complete)
		}()
	}
}

func (e *entry) call(f Func, key string, cancel <-chan struct{}) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, cancel)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}

//!-monitor
