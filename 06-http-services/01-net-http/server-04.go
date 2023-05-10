package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// application specific handlers

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

func productsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers details will be served")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

//middlewares
func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func profileMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		log.Println(elapsed)
	}
}

//utility function for assembling middlewares
type Middleware func(http.HandlerFunc) http.HandlerFunc

func chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func main() {
	// srv := &appServer{}
	/*
		srv := http.DefaultServeMux
		srv.HandleFunc("/", indexHandler)
		srv.HandleFunc("/customers", customersHandler)
		srv.HandleFunc("/products", productsHandler)
	*/
	/*
		indexWithLogHandler := logMiddleware(indexHandler)
		http.HandleFunc("/", indexWithLogHandler)
	*/

	/*
		http.HandleFunc("/", profileMiddleware(logMiddleware(indexHandler)))
		http.HandleFunc("/customers", profileMiddleware(logMiddleware(customersHandler)))
		http.HandleFunc("/products", profileMiddleware(logMiddleware(productsHandler)))
	*/
	http.HandleFunc("/", chain(indexHandler, logMiddleware, profileMiddleware))
	http.HandleFunc("/customers", chain(customersHandler, logMiddleware, profileMiddleware))
	http.HandleFunc("/products", chain(productsHandler, logMiddleware, profileMiddleware))

	http.ListenAndServe(":8080", nil)
}
