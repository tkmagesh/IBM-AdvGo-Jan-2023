package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		LogOperation("Add", 100, 200)
		LogOperation("subtract", 100, 200)
		LogOperation("multiply", 100, 200)
	*/

	/*
		LogOperation(add, 100, 200)
		LogOperation(subtract, 100, 200)
		LogOperation(func(x, y int) {
			fmt.Println("Multiply result :", x*y)
		}, 100, 200)
	*/

	// composing log
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)

		logMultiply := getLogOperation(func(x, y int) {
			fmt.Println("Multiply result :", x*y)
		})
		logMultiply(100, 200)
	*/

	// composing profiling
	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfileOperation(subtract)
		profileSubtract(100, 200)
	*/

	//composing log & profile
	/*
		logAdd := getLogOperation(add)
		profileLogAdd := getProfileOperation(logAdd)
	*/
	profileLogAdd := getProfileOperation(getLogOperation(add))
	profileLogAdd(100, 200)

}

func getProfileOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started...")
		operation(x, y)
		log.Println("Operation completed!")
	}
}

/*
func LogOperation(operation func(int, int), x, y int) {
	log.Println("Operation started...")
	operation(x, y)
	log.Println("Operation completed!")
}
*/

/*
func LogOperation(operationName string, x, y int) {
	switch operationName {
	case "add":
		add(x, y)
	case "subtract":
		subtract(x, y)
	default:
		panic("invalid operation name")
	}
}
*/

/*
func logAdd(x, y int) {
	log.Println("Operation started...")
	add(x, y)
	log.Println("Operation completed!")
}

func logSubtract(x, y int) {
	log.Println("Operation started...")
	subtract(x, y)
	log.Println("Operation completed!")
}
*/

//3rd party library
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
