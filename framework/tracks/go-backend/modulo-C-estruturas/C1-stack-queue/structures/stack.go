package structures

import (
	"errors"
	"sync"
)

type Stack struct {
	items []string
	mu    sync.Mutex
}

var EmptyStack = errors.New("Stack is Empty")

func (s *Stack) Push(item string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}
func (s *Stack) Pop() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.isEmpty() {
		return "", EmptyStack
	}
	lastIndex := len(s.items) - 1
	copyDeletedVal := s.items[lastIndex]
	s.items[lastIndex] = ""
	s.items = s.items[:lastIndex]
	return copyDeletedVal, nil
}

func (s *Stack) Peek() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.isEmpty() {
		return "", EmptyStack
	}
	lastIndex := len(s.items) - 1
	return s.items[lastIndex], nil
}

func (s *Stack) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.isEmpty()
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}
