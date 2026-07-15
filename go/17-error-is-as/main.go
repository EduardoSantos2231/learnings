package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ValidationError struct {
	Field string
}

func main() {
	handleError(processItem(""))
	handleError(processItem("0"))
	handleError(processItem("x"))
	handleError(processItem("42"))
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field: %s", e.Field)
}

func processItem(id string) error {
	switch id {
	case "":
		return ErrNotFound
	case "0":
		return &ValidationError{
			Field: "id",
		}
	case "x":
		return fmt.Errorf("db error %w", ErrNotFound)
	default:
		return nil
	}
}

func handleError(err error) {
	if err == nil {
		return
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println(err.Error())
		return
	}
	if target, ok := errors.AsType[*ValidationError](err); ok {
		fmt.Println("Field not valid: ", target.Field)
		return
	}
}
