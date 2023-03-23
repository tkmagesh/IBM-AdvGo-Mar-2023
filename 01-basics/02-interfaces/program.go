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

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

// (Area) step-1
/*
func printArea(x interface{}) {
	if shapeWithArea, ok := x.(interface{ Area() float32 }); ok {
		fmt.Printf("Area = %0.2f\n", shapeWithArea.Area())
	} else {
		fmt.Println("Given object does not have the Area() method")
	}
}
*/

// (Area) step-2
/*
func printArea(x interface{ Area() float32 }) {
	fmt.Printf("Area = %0.2f\n", x.Area())
}
*/

// (Area) step-3
type AreaFinder interface {
	Area() float32
}

func printArea(x AreaFinder) {
	fmt.Printf("Area = %0.2f\n", x.Area())
}

// (Perimeter)

type PerimeterFinder interface {
	Perimeter() float32
}

func printPerimeter(x PerimeterFinder) {
	fmt.Printf("Perimeter = %0.2f\n", x.Perimeter())
}

//
/*
func printShapeStats(x interface {
	Area() float32
	Perimeter() float32
}) {
	printArea(x)      // x has to be interface { Area() float32 }
	printPerimeter(x) // x has to interface { Perimeter() float32 }
}
*/

/*
func printShapeStats(x interface {
	interface {
		Area() float32
	} // => AreaFinder
	interface {
		Perimeter() float32
	} // => PerimeterFinder
}) {
	printArea(x)      // x has to be interface{ Area() float32 }
	printPerimeter(x) // x has to interface { Perimeter() float32 }
}
*/

/*
func printShapeStats(x interface {
	AreaFinder
	PerimeterFinder
}) {
	printArea(x)      // x has to be interface{ Area() float32 }
	printPerimeter(x) // x has to interface { Perimeter() float32 }
}
*/

type Shape interface {
	AreaFinder
	PerimeterFinder
}

func printShapeStats(x Shape) {
	printArea(x)      // x has to be interface{ Area() float32 }
	printPerimeter(x) // x has to interface { Perimeter() float32 }
}

// do the above for Perimeter (Circle = 2 * pi * r, Rectangle = 2 * (Height + Width))
func main() {

	// printArea(100)

	c := Circle{Radius: 12}
	// fmt.Printf("Area = %0.2f\n", c.Area())
	/*
		printArea(c)
		printPerimeter(c)
	*/
	printShapeStats(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Printf("Area = %0.2f\n", r.Area())
	/*
		printArea(r)
		printPerimeter(r)
	*/
	printShapeStats(r)
}
