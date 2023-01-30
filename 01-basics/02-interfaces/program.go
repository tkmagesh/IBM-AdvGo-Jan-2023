/*
Interfaces
	- contracts
	- IMPLICITLY implemented
*/
package main

import (
	"fmt"
	"math"
)

/* 3rd party library */
type Circle struct /* implements AreaFinder, PerimeterFinder, ShapeStatsFinder */ {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Width)
}

/* our app */
/* Utility functions */
type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/*
func PrintPerimeter(x interface {
	Perimeter() float32
}) {
	fmt.Println("Perimeter :", x.Perimeter())
}
*/

/* ----- */
/* interface composition */

/*
func PrintShapeStats(x interface {
	AreaFinder
	PerimeterFinder
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

/*
func PrintShapeStats(x interface {
	Area() float32
	Perimeter() float32
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShapeStats(x ShapeStatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

/* struct composition */
type Square struct {
	Rectangle
}

func NewSquare(sideLength float32) *Square {
	return &Square{
		Rectangle: Rectangle{Length: sideLength, Width: sideLength},
	}
}

func main() {
	c := Circle{Radius: 12}
	/*
		// fmt.Println("Area :", c.Area())
		PrintArea(c)
		// fmt.Println("Perimeter :", c.Perimeter())
		PrintPerimeter(c)
	*/
	PrintShapeStats(c)

	r := Rectangle{Length: 10, Width: 12}
	/*
		// fmt.Println("Area :", r.Area())
		PrintArea(r)
		// fmt.Println("Perimeter :", r.Perimeter())
		PrintPerimeter(r)
	*/
	PrintShapeStats(r)

	/* inheritence through composition */
	s := NewSquare(10)
	/*
		PrintArea(s)
		PrintPerimeter(s)
	*/
	PrintShapeStats(s)

}
