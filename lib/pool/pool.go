package pool

import (
	"errors"
	"fmt"
	"pca/lib/jlog"
	"runtime"
	"sync"
	"time"
)

type (
	Src interface {
		Usable() bool
		Reset()
		Close()
	}

	Pool interface {
		Call(func(Src) error) error
		Close()
		Len() int
	}

	Factory func() (Src, error)

	classic struct {
		srcs     chan Src
		capacity int
		maxIdle  int
		len      int
		factory  Factory
		gcTime   time.Duration
		closed   bool
		sync.RWMutex
	}
)

func ClassicPool(capacity, maxIdle int,
	factory Factory, gcTime time.Duration) Pool {
	pool := &classic{
		srcs:     make(chan Src, capacity),
		capacity: capacity,
		maxIdle:  maxIdle,
		factory:  factory,
		gcTime:   gcTime,
		closed:   false,
	}
	go pool.gc()
	return pool
}

func (self *classic) Call(callback func(Src) error) error {
	var src Src
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("%v", p)
			jlog.LogError(err)
		}
		self.recover(src)
	}()

	for {
		self.RLock()
		if self.closed {
			self.RUnlock()
			return errors.New("资源池已关闭")
		}
		select {
		case src = <-self.srcs:
			self.RUnlock()
			if !src.Usable() {
				self.del(src)
				break
			}
		default:
			self.RUnlock()
			err := self.incAuto()
			if err != nil {
				fmt.Print(err.Error())
				return err
			}
			runtime.Gosched()
			continue
		}

		return callback(src)
	}

}

func (self *classic) Close() {
	self.Lock()
	defer self.Unlock()

	if self.closed {
		return
	}
	self.closed = true
	for i := len(self.srcs); i >= 0; i-- {
		(<-self.srcs).Close()
	}
	close(self.srcs)
	self.len = 0
}

func (self *classic) Len() int {
	self.RLock()
	defer self.RUnlock()
	return self.len
}

func (self *classic) gc() {
	for !self.isClosed() {
		self.Lock()
		extra := len(self.srcs) - self.maxIdle
		if extra > 0 {

			self.len -= extra
			for ; extra > 0; extra-- {
				(<-self.srcs).Close()
			}
		}
		self.Unlock()
		time.Sleep(self.gcTime)
	}
}

func (self *classic) incAuto() error {
	self.Lock()
	defer self.Unlock()
	if self.len >= self.capacity {
		return nil
	}
	src, err := self.factory()
	if err != nil {
		return err
	}
	self.srcs <- src
	self.len++
	return nil
}

func (self *classic) del(src Src) {
	src.Close()
	self.Lock()
	self.len--
	self.Unlock()
}

func (self *classic) recover(src Src) {
	self.RLock()
	defer self.RUnlock()
	if self.closed {
		return
	}
	src.Reset()
	self.srcs <- src
}

func (self *classic) isClosed() bool {
	self.RLock()
	defer self.RUnlock()
	return self.closed
}
