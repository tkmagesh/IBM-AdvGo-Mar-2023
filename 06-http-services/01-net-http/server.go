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
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
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
