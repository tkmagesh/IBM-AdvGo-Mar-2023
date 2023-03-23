package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

// step-1
/*
func printArea(x interface{}) {
	if shapeWithArea, ok := x.(interface{ Area() float32 }); ok {
		fmt.Printf("Area = %0.2f\n", shapeWithArea.Area())
	} else {
		fmt.Println("Given object does not have the Area() method")
	}
}
*/

// step-2
/*
func printArea(x interface{ Area() float32 }) {
	fmt.Printf("Area = %0.2f\n", x.Area())
}
*/

// step-3
type AreaFinder interface {
	Area() float32
}

func printArea(x AreaFinder) {
	fmt.Printf("Area = %0.2f\n", x.Area())
}

// do the above for Perimeter (Circle = 2 * pi * r, Rectangle = 2 * (Height + Width))
func main() {

	// printArea(100)

	c := Circle{Radius: 12}
	// fmt.Printf("Area = %0.2f\n", c.Area())
	printArea(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Printf("Area = %0.2f\n", r.Area())
	printArea(r)
}
