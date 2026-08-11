package storage

import (
	"errors"
	"fmt"
	"sync"
)

var ErrIdNotFound = errors.New("Id not found")

type Item struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Done bool   `json:"done"`
}

type storage struct {
	Items []Item
	mu    sync.Mutex
}

func NewStorage() *storage {
	return &storage{
		Items: []Item{},
	}
}

func (s *storage) AddItem(name string) Item {
	s.mu.Lock()
	defer s.mu.Unlock()
	lasIn := len(s.Items)
	newItem := Item{
		Name: name,
		ID:   lasIn + 1,
	}
	s.Items = append(s.Items, newItem)
	return newItem
}

func (s *storage) ListAllItems() []Item {
	s.mu.Lock()
	defer s.mu.Unlock()
	cpy := make([]Item, len(s.Items), cap(s.Items))
	copy(cpy, s.Items)
	return cpy
}

func (s *storage) MarkAsDone(id ...int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range id {
		found := s.auxSetDoneIdBased(v)
		if !found {
			return ErrIdNotFound
		}
	}
	return nil
}

func (s *storage) auxSetDoneIdBased(id int) bool {
	found := false
	for index := range s.Items {
		if s.Items[index].ID == id {
			s.Items[index].Done = true
			found = true
		}
	}
	return found
}

func (s *storage) Seed() {
	for i := range 10 {
		newName := fmt.Sprintf("Test %d", i)
		s.AddItem(newName)
	}
}

func (s *storage) ResetDB() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Items = []Item{}
}
