package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"price"`
}

var products = []Product{
	{Id: 101, Name: "Pen", Cost: 10},
	{Id: 102, Name: "Pencil", Cost: 5},
	{Id: 103, Name: "Marker", Cost: 50},
}

type appServer struct {
}

//http.Handler interface implementation
func (server *appServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "invalid request", http.StatusBadRequest)
			}
			newProduct.Id = len(products) + 101
			products = append(products, newProduct)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		default:
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "All the customers details will be served")
	case "/":
		fmt.Fprintln(w, "Hello World!")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}

}

func main() {
	srv := &appServer{}
	http.ListenAndServe(":8080", srv)
}
