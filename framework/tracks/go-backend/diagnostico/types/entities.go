// package types makes shure that we are using known types for our domain
package types

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
}
