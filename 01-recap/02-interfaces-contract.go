package main

import (
	"fmt"
	"math"
)

// Sprint-1
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Sprint-2
type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

/*
func PrintArea(x interface{}) {
	// fmt.Println("Area :", x.Area())
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	case interface{}:
		fmt.Println("x has no Area() method")
	}
}
*/

/*
func PrintArea(x interface{}) {
	// fmt.Println("Area :", x.Area())
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	case interface{}:
		fmt.Println("x has no Area() method")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

// Sprint-3
// Calculate Perimeter for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

func PrintPerimeter(x interface{ Perimeter() float64 }) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// Sprint-4
func PrintStats(x interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface{ Perimeter() float64 }
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{Length: 10, Breadth: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)
}
