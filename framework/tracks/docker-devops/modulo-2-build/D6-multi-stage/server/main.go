// server that says hello world
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from multi-stage build!")
	})
	http.ListenAndServe(":8080", nil)
}
