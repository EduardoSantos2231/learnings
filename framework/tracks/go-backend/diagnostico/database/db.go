// package database will emulate a database and take care of the CRUD
package database

import (
	"diagnostico/types"
	"sync"
)

// TO-DO:
// Implement these methods
type Storer interface {
	Insert(t types.Task)
	Delete(tID string)
	Update(t types.Task)
}

type database struct {
	source []types.Task
	mux    sync.RWMutex
}

func NewDatabase() *database {
	source := make([]types.Task, 1)
	return &database{
		source: source,
	}
}
