/*
	- Assign functions to variables
	- Pass functions are arguments
	- Return functions as return values
*/

package main

import "fmt"

func main() {
	/*
		var fn func()
		fn = func() {
			fmt.Println("fn invoked")
		}
		fn()
	*/
	var fx func()
	fx = fn
	fx()

	fx = greet
	fx()

	var operation func(x, y int)

	operation = func(x, y int) {
		fmt.Println("Add Result :", x+y)
	}
	operation(100, 200)

	operation = func(x, y int) {
		fmt.Println("Subtract Result :", x-y)
	}
	operation(100, 200)
}

func fn() {
	fmt.Println("fn invoked")
}

func greet() {
	fmt.Println("Hi there!")
}
