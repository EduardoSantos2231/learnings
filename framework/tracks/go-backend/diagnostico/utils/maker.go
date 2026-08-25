// package maker will make shure that we are providin valid data to our database
package maker

import (
	"diagnostico/types"
)

func CreateTask(name, description string) *types.Task {
	id := genRandomId()

	return &types.Task{
		Name:        name,
		Description: description,
		ID:          id,
	}
}

func genRandomId() string {
	// think in a form to gen a random id
}
