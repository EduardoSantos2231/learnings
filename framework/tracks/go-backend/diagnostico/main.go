package main

import (
	"diagnostico/database"
	"net/http"
)

// TO-DO:
// make a healthCheck Route
// test it
//
// General structure idea:
// maker will give us the tools to make a Task
// types will provide types do make a standard way for us to take care of requests and responses
// utils will give us tools to optmize the way that we write our responses and encode them...

var db = database.NewDatabase()

func main() {
	mux := http.NewServeMux()
	wrppdMux := jsonMiddleware(mux)

	http.ListenAndServe(":8080", wrppdMux)
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application-json")
		next.ServeHTTP(w, r)
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}
