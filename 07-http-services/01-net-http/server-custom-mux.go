package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/* custom mux */
type WebApp struct {
	handlers map[string]map[string]http.HandlerFunc
}

func (wa *WebApp) Register(method string, pattern string, handler http.HandlerFunc) {
	if _, exists := wa.handlers[method]; !exists {
		wa.handlers[method] = make(map[string]http.HandlerFunc)
	}
	wa.handlers[method][pattern] = handler
}

func (wa *WebApp) Get(pattern string, handler http.HandlerFunc) {
	wa.Register("GET", pattern, handler)
}

func (wa *WebApp) Post(pattern string, handler http.HandlerFunc) {
	wa.Register("POST", pattern, handler)
}

func (wa *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodHandlers, exists := wa.handlers[r.Method]
	if !exists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if handler, exists := methodHandlers[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Requested resource not found")
}

func NewWebApp() *WebApp {
	return &WebApp{
		handlers: make(map[string]map[string]http.HandlerFunc),
	}
}

// handlers
func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All the product info will be served")
}

func newProductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The given new product will be added")
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All the customer requests will be processed")
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All the order requests will be processed")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
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
	// app := &WebApp{}
	app := NewWebApp()
	app.Get("/products", chain(productsHandler, logMiddleware, profileMiddleware))
	app.Post("/products", chain(newProductHandler, logMiddleware, profileMiddleware))
	app.Get("/customers", chain(customersHandler, logMiddleware, profileMiddleware))
	app.Get("/orders", chain(ordersHandler, logMiddleware, profileMiddleware))
	app.Get("/", chain(indexHandler, logMiddleware, profileMiddleware))

	/*
		app.Register("GET", "/products", productsHandler)
		app.Register("POST", "/products", newProductHandler)
		app.Register("GET", "/customers", customersHandler)
		app.Register("GET", "/orders", ordersHandler)
		app.Register("GET", "/", indexHandler)
	*/

	http.ListenAndServe(":8080", app)
}
