package main

import (
	"fmt"
	"log"
	"net/http"
)

type WebApp struct {
}

func (wa *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/products":
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "All the product info will be served")
		case "POST":
			fmt.Fprintf(w, "The given new product will be added")
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "Method not allowed")
		}

	case "/customers":
		fmt.Fprintf(w, "All the customer requests will be processed")
	case "/orders":
		fmt.Fprintf(w, "All the order requests will be processed")
	case "/":
		fmt.Fprintln(w, "Hello World!")
	default:
		w.WriteHeader(404)
		fmt.Fprint(w, "Resource not found!")
	}

}

func main() {
	app := &WebApp{}
	http.ListenAndServe(":8080", app)
}
