package main

import (
	"fmt"
	"log"
	"net/http"
)

/* custom mux */
type WebApp struct {
	handlers map[string]http.HandlerFunc
}

func (wa *WebApp) Register(pattern string, handler http.HandlerFunc) {
	wa.handlers[pattern] = handler
}

func (wa *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	if handler, exists := wa.handlers[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Requested resource not found")
}

func NewWebApp() *WebApp {
	return &WebApp{
		handlers: make(map[string]http.HandlerFunc),
	}
}

// handlers
func productsHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, "All the customer requests will be processed")
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All the order requests will be processed")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	// app := &WebApp{}
	app := NewWebApp()
	app.Register("/products", productsHandler)
	app.Register("/customers", customersHandler)
	app.Register("/orders", ordersHandler)
	app.Register("/", indexHandler)
	http.ListenAndServe(":8080", app)
}
