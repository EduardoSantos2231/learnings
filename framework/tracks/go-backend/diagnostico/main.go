package main

import (
	"context"
	"diagnostico/data_types"
	"diagnostico/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/tools/go/analysis/passes/defers"
)

var FakeDb = storage.NewStorage()

type validator interface {
	isValid() bool
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/items", handleItemAddition)
	mux.HandleFunc("GET /api/v1/items", handleItemListing)
	mux.HandleFunc("PUT /api/v1/items", handleSetItemDone)
	wrappedMux := LogMiddleware(JSONMiddleware(mux))
	fmt.Println("Listening at http://localhost:3000")
	http.ListenAndServe(":3000", wrappedMux)
}

// METHOD: POST
// Action: add an item to the storage
func handleItemAddition(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var uD datatypes.UserData
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(&uD)
	if err != nil {
		setStatus(w, http.StatusBadRequest)
		return
	}
	if !uD.IsValid() {
		setStatus(w, http.StatusBadRequest)
		return
	}
	itemAdded := FakeDb.AddItem(uD.Name)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		Item_added storage.Item
	}{
		Item_added: itemAdded,
	})
}

// METHOD: GET
// Action: list all the items in FakeDb
func handleItemListing(w http.ResponseWriter, _ *http.Request) {
	allItems := FakeDb.ListAllItems()
	response := struct {
		Items []storage.Item `json:"items"`
	}{
		Items: allItems,
	}
	json.NewEncoder(w).Encode(response)
}

func handleSetItemDone(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userReq datatypes.UserReqData
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		setStatus(w, http.StatusBadRequest)
		return
	}
	if !userReq.IsValid() {
		setStatus(w, http.StatusBadRequest)
		return
	}
	err = FakeDb.MarkAsDone(userReq.Ids...)
	if err != nil {
		setStatus(w, http.StatusBadRequest)
		return
	}
	setStatus(w, http.StatusOK)
}

// Auxiliary function to set headers in response
func setStatus(w http.ResponseWriter, errCode int) {
	w.WriteHeader(errCode)
}

// Auxiliary functiona that runs inside the LogMiddleware
func provideDetails(_ http.ResponseWriter, r *http.Request) {
	log.Println("------------------------------------")
	log.Println("Route: ", r.URL.Path)
	log.Println("Method: ", r.Method)
	log.Println("Client IP: ", r.RemoteAddr)
	log.Println("------------------------------------")
}

// Middleware that ensures the JSON format when answering the request
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Middleware that ensures we know some infos about the client
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		provideDetails(w, r)
		next.ServeHTTP(w, r)
	})
}

func chainValidation(r *http.Request, v ...validator) {
	limit := make(chan struct{}, 3)
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	for _, val := range v {
		go func(val validator) {

		}(val)
	}
}
