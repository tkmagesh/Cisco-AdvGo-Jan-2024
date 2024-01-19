package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float32 `json:"cost"`
	Category string  `json:"category"`
}

var products = []Product{
	{Id: 100, Name: "Pen", Cost: 10, Category: "Stationary"},
	{Id: 101, Name: "Pencil", Cost: 5, Category: "Stationary"},
	{Id: 102, Name: "Marker", Cost: 50, Category: "Stationary"},
	{Id: 103, Name: "Mouse", Cost: 250, Category: "Electronics"},
}

// App specific handlers
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "error encoding data", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "Invalid data format", http.StatusBadRequest)
		}
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All customer data will be served")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/products", ProductsHandler)
	http.HandleFunc("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}
