package main

import "fmt"

func sumInt(list []int) int {
	var result int
	for _, no := range list {
		result += no
	}
	return result
}

func sumFloat32(list []float32) float32 {
	var result float32
	for _, no := range list {
		result += no
	}
	return result
}

//generics
type Numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// func sum[T int | float32](list []T) T {
func sum[T Numbers](list []T) T {
	var result T
	for _, no := range list {
		result += no
	}
	return result
}

func main() {
	ints := []int{3, 1, 4, 2, 5}
	// fmt.Println(sumInt(ints))
	sumInt := sum(ints)
	fmt.Println(sumInt)

	floats := []float32{3, 1.8, 4.3, 2.8, 5.7}
	sumFloat := sum(floats)
	fmt.Println(sumFloat)
}
