/*
	Need for pointers in Go?
		EVERYTHING IS A VALUE in GO
			Assignment operation "Creates" a copy
			Comparison is ALWAYS by value
*/
package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

func (emp *Employee) ChangeName(newName string) {
	fmt.Printf("[ChangeName] emp = %p\n", emp)
	emp.Name = newName
}

func main() {
	/* Arrays (values) */
	n1 := [3]int{10, 20, 30}
	n2 := [3]int{10, 20, 30}
	fmt.Printf("&n1 = %p, &n2 = %p\n", &n1, &n2)
	// comparison is by value
	fmt.Println("n1 == n2 ? :", n1 == n2)

	//Assignment operation results in creating a "copy"
	n3 := n1
	n3[0] = 100
	fmt.Printf("n3 = %v, n1 = %v \n", n3, n1)

	/* Structs (values) */
	e1 := Employee{Id: 100, Name: "Magesh"}
	e2 := Employee{Id: 100, Name: "Magesh"}
	fmt.Printf("&e1 = %p, &e2 = %p\n", &e1, &e2)
	fmt.Println("e1 == e2 ? :", e1 == e2)

	/* Slices (reference) */
	s1 := []int{10, 20, 30}
	s2 := s1
	s2[0] = 100
	fmt.Printf("s1 = %v, s2 = %v \n", s1, s2)

	/* Pointers */

	var x int
	x = 100

	var xPtr *int
	xPtr = &x // (value => address)
	fmt.Println(xPtr, x)

	// dereferencing (address => value)
	fmt.Println(*xPtr)

	change(&x)
	fmt.Println(x)

	// Pointers in struct methods
	fmt.Printf("[main] &e1 = %p\n", &e1)
	fmt.Println("Before changing the name, e1 = ", e1)
	// (&e1).ChangeName("Suresh")
	e1.ChangeName("Suresh")
	fmt.Println("After changing the name, e1 = ", e1)
}

func change(n *int) {
	*n = *n * 10
}
