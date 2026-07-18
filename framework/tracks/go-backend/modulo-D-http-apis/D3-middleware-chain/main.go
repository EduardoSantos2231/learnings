package main

import (
	"encoding/json"
	"fmt"
	"log"
	"middleware-chain/product"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type middleware func(http.Handler) http.Handler

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
	middlewares := []middleware{
		recovererMiddleware,
		loggerMiddleware,
		jsonMiddleware,
		authMiddleware,
		corsMiddleware,
	}
	finalHandler := chain(mux, middlewares...)
	fmt.Println("server has started at: http://localhost:8080")
	err := http.ListenAndServe(":8080", finalHandler)
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
		dealError(w, http.StatusBadRequest, err.Error())
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
		dealError(w, http.StatusBadRequest, err.Error())
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
		dealError(w, http.StatusBadRequest, err.Error())
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
		dealError(w, http.StatusBadRequest, err.Error())
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

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		duration := time.Now()
		log.Println("[METHOD]: ", r.Method)
		log.Println("[ROUTE PATH]: ", r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("[DURATION]: ", time.Since(duration))
	})
}

func chain(finalHandler http.Handler, middlewares ...middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		finalHandler = middlewares[i](finalHandler)
	}
	return finalHandler
}

func recovererMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("New error ocurred and crashed the server: ", err)
				writer := json.NewEncoder(w)
				w.WriteHeader(http.StatusInternalServerError)
				writer.Encode(struct {
					Message string
				}{
					Message: "something unexpected happened, we are very sorry",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})

}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/products" {
			next.ServeHTTP(w, r)
			return
		}
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		bearerTkn := strings.Split(token, " ")
		if bearerTkn[1] != "admin123" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Methods", "PUT")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
