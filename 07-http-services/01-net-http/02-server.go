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

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MyServer struct {
	routes      map[string]http.HandlerFunc
	middlewares []Middleware
}

func (myServer *MyServer) Register(pattern string, handler http.HandlerFunc) {
	// myServer.routes[pattern] = handler
	// to ensure that the handler func is executed after executing the middlewares
	for _, middleware := range myServer.middlewares {
		handler = middleware(handler)
	}
	myServer.routes[pattern] = handler
}

func (myServer *MyServer) UseMiddleware(middleware Middleware) {
	myServer.middlewares = append(myServer.middlewares, middleware)
}

// http.Handler interface implementation
func (myServer *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exists := myServer.routes[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

func NewMyServer() *MyServer {
	return &MyServer{
		routes: make(map[string]http.HandlerFunc),
	}
}

// handler functions (type : http.HandlerFunc)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%s - %s\n", r.Method, r.URL.Path) // refactored as a middleware
	fmt.Fprintln(w, "Hello World!")
}

func productsHanlder(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%s - %s\n", r.Method, r.URL.Path) // refactored as a middleware
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
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%s - %s\n", r.Method, r.URL.Path) // refactored as a middleware
	fmt.Fprintln(w, "All customer requests will be processed")
}

// middlewares
func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func main() {
	server := NewMyServer()
	// using middlewares
	server.UseMiddleware(logMiddleware)
	server.Register("/", indexHandler)
	server.Register("/products", productsHanlder)
	server.Register("/customers", customersHandler)
	http.ListenAndServe(":8080", server)
}
