package product

import (
	"errors"
	"sync"
)

var ErrProductNotFound = errors.New("Product Not Found")

type Store struct {
	items map[int]Product
	mu    sync.RWMutex
}

func (s *Store) List() []Product {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []Product
	for _, el := range s.items {
		result = append(result, el)
	}
	return result
}

func (s *Store) GetById(id int) (Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	product, ok := s.items[id]
	if !ok {
		return Product{}, ErrProductNotFound
	}
	return product, nil
}

func (s *Store) Create(name string, price float64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	prod := NewProduct(name, price)
	s.items[prod.ID] = *prod
	return prod.ID, nil
}

func (s *Store) Update(id int, updatedProduct Product) (Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.items[id]
	if !ok {
		return Product{}, ErrProductNotFound
	}
	s.items[id] = updatedProduct
	return s.items[id], nil
}

func (s *Store) Delete(id int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.items[id]
	if !ok {
		return 0, ErrProductNotFound
	}
	delete(s.items, id)
	return id, nil
}

func NewStore() *Store {
	return &Store{
		items: make(map[int]Product),
	}
}
