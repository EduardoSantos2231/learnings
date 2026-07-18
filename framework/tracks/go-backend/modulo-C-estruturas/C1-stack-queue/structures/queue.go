package structures

import (
	"errors"
	"sync"
)

type Queue struct {
	items []string
	mu    sync.Mutex
}

var EmptyQueue error = errors.New("Empty Queue")

func (q *Queue) Enqueue(item string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (string, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isEmpty() {
		return "", EmptyQueue
	}
	deletedVal := q.items[0]
	q.items = q.items[1:]
	return deletedVal, nil
}

func (q *Queue) Peek() (string, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isEmpty() {
		return "", EmptyQueue
	}
	return q.items[0], nil
}

func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.isEmpty()
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}
