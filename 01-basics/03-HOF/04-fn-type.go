package main

import (
	"fmt"
	"log"
	"time"
)

type OperationFn func(int, int)

func main() {

	profileLogAdd := getProfileOperation(getLogOperation(add))
	profileLogAdd(100, 200)

	getProfileOperation(getLogOperation(subtract))(100, 200)
}

func getProfileOperation(operation OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

func getLogOperation(operation OperationFn) OperationFn {
	return func(x, y int) {
		log.Println("Operation started...")
		operation(x, y)
		log.Println("Operation completed!")
	}
}

//3rd party library
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
