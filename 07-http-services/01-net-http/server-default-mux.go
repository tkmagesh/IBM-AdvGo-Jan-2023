package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// handlers
func productsHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "All the product info will be served")
	case "POST":
		fmt.Fprintf(w, "The given new product will be added")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")
	}
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("%s - %s\n", r.Method, r.URL.Path)
	fmt.Fprintf(w, "All the customer requests will be processed")
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("%s - %s\n", r.Method, r.URL.Path)
	fmt.Fprintf(w, "All the order requests will be processed")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("%s - %s\n", r.Method, r.URL.Path)
	fmt.Fprintln(w, "Hello World!")
}

// middlewares
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
		fmt.Printf("time : %v\n", elapsed)
	}
}

// utility
type Middleware func(http.HandlerFunc) http.HandlerFunc

func chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func main() {
	/*
		http.HandleFunc("/products", profileMiddleware(logMiddleware(productsHandler)))
		http.HandleFunc("/customers", profileMiddleware(logMiddleware(customersHandler)))
		http.HandleFunc("/orders", profileMiddleware(logMiddleware(ordersHandler)))
		http.HandleFunc("/", profileMiddleware(logMiddleware(indexHandler)))
	*/
	http.HandleFunc("/products", chain(productsHandler, logMiddleware, profileMiddleware))
	http.HandleFunc("/customers", chain(customersHandler, logMiddleware, profileMiddleware))
	http.HandleFunc("/orders", chain(ordersHandler, logMiddleware, profileMiddleware))
	http.HandleFunc("/", chain(indexHandler, logMiddleware, profileMiddleware))
	http.ListenAndServe(":8080", nil)
}
