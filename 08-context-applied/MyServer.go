package main

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MyServer struct {
	routes      map[string]http.HandlerFunc
	middlewares []Middleware
}

func (myserver *MyServer) UseMiddleware(middleware Middleware) {
	myserver.middlewares = append(myserver.middlewares, middleware)
}

func (myserver *MyServer) AddRoute(pattern string, handler http.HandlerFunc) {
	for _, middleware := range myserver.middlewares {
		handler = middleware(handler)
	}
	myserver.routes[pattern] = handler
}

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
