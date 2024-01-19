package main

import (
	"context"
	"context-app/customers"
	"context-app/index"
	"context-app/products"
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/google/uuid"
)

func LoggerMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%s - %s", r.Method, r.URL.Path)
		start := time.Now()
		handler(w, r)
		// statusCode := w.
		elapsed := time.Since(start)
		log.Printf("%s %v\n", msg, elapsed)
	}
}

func TraceIdMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := uuid.New()
		valCtx := context.WithValue(r.Context(), "request-id", reqId)
		handler(w, r.WithContext(valCtx))
	}
}

func main() {
	// log.SetFormatter(&log.JSONFormatter{})

	myServer := NewMyServer()
	myServer.UseMiddleware(LoggerMiddleware)
	myServer.UseMiddleware(TraceIdMiddleware)

	myServer.AddRoute("/", index.Handler)
	myServer.AddRoute("/customers", customers.Handler)
	myServer.AddRoute("/products", products.Handler)
	if err := http.ListenAndServe(":8080", myServer); err != nil {
		log.Println(err)
		return
	}
}
