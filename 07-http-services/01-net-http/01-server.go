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
	Units    int     `json:"-"`
	Category string  `json:"category"`
}

var products = []Product{
	{101, "Pen", 10, 20, "Stationary"},
	{102, "Pencil", 5, 200, "Stationary"},
	{103, "Marker", 50, 10, "Stationary"},
}

type MyServer struct {
}

// http.Handler interface implementation
func (server *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		// fmt.Fprintln(w, "All product requests will processed")
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "input data error", http.StatusBadRequest)
				return
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
		}
		/*

		 */
		// Assumption : a POST request is made

	case "/customers":
		fmt.Fprintln(w, "All customer requests will be processed")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}

}

func main() {
	server := &MyServer{}
	http.ListenAndServe(":8080", server)
}
