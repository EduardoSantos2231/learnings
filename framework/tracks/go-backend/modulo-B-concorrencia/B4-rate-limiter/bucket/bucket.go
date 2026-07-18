package bucket

import (
	"sync"
	"time"
)

type Bucket struct {
	tokens   int
	capacity int
	interval time.Duration
	mu       sync.Mutex
	stopChan chan bool
}

func NewBucket(capacity int, interval time.Duration) *Bucket {
	stopChan := make(chan bool, 1)
	result := &Bucket{
		tokens:   capacity,
		capacity: capacity,
		interval: interval,
		stopChan: stopChan,
	}

	go result.increaseTokens()

	return result
}

func (b *Bucket) increaseTokens() {
	timer := time.NewTicker(b.interval)
	defer timer.Stop()
	for {
		select {
		case <-b.stopChan:
			return
		case <-timer.C:
			b.mu.Lock()
			if b.tokens < b.capacity {
				b.tokens++
			}
			b.mu.Unlock()
		}
	}
}

func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.tokens == 0 {
		return false
	}
	b.tokens--
	return true
}

func (b *Bucket) Stop() {
	b.stopChan <- true
}
