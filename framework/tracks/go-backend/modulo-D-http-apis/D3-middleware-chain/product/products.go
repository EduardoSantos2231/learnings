package product

import "sync"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type id struct {
	id int
	mu sync.RWMutex
}

var globalID = id{
	id: 0,
}

func NewProduct(name string, price float64) *Product {

	return &Product{
		ID:    globalID.genIncreasingId(),
		Name:  name,
		Price: price,
	}
}

func (i *id) genIncreasingId() int {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.id++
	return i.id
}
