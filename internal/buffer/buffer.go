package buffer

import (
	"sync"
	"time"
)

// Buffer is used to store the buffer data of trade volume and is pushed to the database only when the
// buffer is full or after certain time interval
type Buffer struct {
	value       float64
	currSize    int64
	maxSize     int64
	lastFlushed time.Time
	lock        sync.Mutex
	funcOnFlush func()
	maxDuration time.Duration
}

func NewBuffer(value float64, size int64, funcOnFlush func()) *Buffer {
	buf := &Buffer{
		value:       value,
		currSize:    size,
		lastFlushed: time.Now(),
		funcOnFlush: funcOnFlush,
		maxSize:     10,
		maxDuration: 5 * time.Second,
	}
	go buf.startFlushTicker()
	return buf
}

// Add adds the value to the buffer and flushes the buffer if the buffer is full or after certain time interval
func (b *Buffer) Add(value float64) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.value += value
	b.currSize++

	if b.currSize >= b.maxSize || time.Since(b.lastFlushed) >= b.maxDuration {
		b.flush()
	}
}

// flush flushes the buffer, calls the callback and resets the buffer
func (b *Buffer) flush() {
	b.funcOnFlush()
	b.value = 0
	b.currSize = 0
	b.lastFlushed = time.Now()
}

// startFlushTicker starts a ticker to flush the buffer after certain time interval
func (b *Buffer) startFlushTicker() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			b.lock.Lock()
			if time.Since(b.lastFlushed) >= b.maxDuration && b.currSize > 0 {
				b.flush()
			}
			b.lock.Unlock()
		}
	}
}
