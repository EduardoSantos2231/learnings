package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"products-api/product"
	"strconv"
)

type database interface {
	List() []product.Product
	GetById(id int) (product.Product, error)
	Create(name string, price float64) (int, error)
	Update(id int, updatedProduct product.Product) (product.Product, error)
	Delete(id int) (int, error)
}

var ourDatabase database = product.NewStore()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", getProducts)
	mux.HandleFunc("POST /products", createProducts)
	mux.HandleFunc("GET /products/{id}", getProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)
	fmt.Println("server has started at: http://localhost:8080")
	err := http.ListenAndServe(":8080", jsonMiddleware(mux))
	if err != nil {
		fmt.Println("deu ruim pae")
		return
	}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	responseReturn := json.NewEncoder(w)
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		dealError(w, http.StatusInternalServerError, err.Error())
		return
	}
	queryResult, err := ourDatabase.GetById(convertedId)
	if err != nil {
		dealError(w, http.StatusNotFound, err.Error())
		return
	}
	responseReturn.Encode(queryResult)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	products := ourDatabase.List()
	response := struct {
		Data []product.Product
	}{
		Data: products,
	}
	encoder.Encode(response)
}

func createProducts(w http.ResponseWriter, r *http.Request) {
	responseWriter := json.NewEncoder(w)
	userData := struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		dealError(w, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := ourDatabase.Create(userData.Name, userData.Price)
	if err != nil {
		dealError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	responseWriter.Encode(struct {
		Message string
		Id      int
	}{
		Message: "created with success",
		Id:      id,
	})
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	userData := struct {
		product.Product
	}{}
	responseWriter := json.NewEncoder(w)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		dealError(w, http.StatusInternalServerError, err.Error())
		return
	}
	updatedProduct, err := ourDatabase.Update(userData.Product.ID, userData.Product)
	if err != nil {
		dealError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(200)
	responseWriter.Encode(struct {
		Message        string
		UpdatedProduct product.Product
	}{
		Message:        "item updated with success",
		UpdatedProduct: updatedProduct,
	})
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	responseWriter := json.NewEncoder(w)
	id := r.PathValue("id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		dealError(w, http.StatusInternalServerError, err.Error())
		return
	}
	deltedID, err := ourDatabase.Delete(convertedId)
	if err != nil {
		dealError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	responseWriter.Encode(struct {
		Message   string
		IdDeleted int
	}{
		Message:   "item deleted",
		IdDeleted: deltedID,
	})
}

func dealError(w http.ResponseWriter, statusCode int, err string) {
	writer := json.NewEncoder(w)
	response := struct {
		Message string
		Details string
	}{
		Message: "something went wrong while executing ur request",
		Details: err,
	}
	w.WriteHeader(statusCode)
	writer.Encode(response)
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
