package models

/* //go:generate go run ../col-gen.go -N Product -P models */
//go:generate col-gen -N Product -P models
//go:generate gofmt -w .
type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}
