package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

// fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) String() string {
	var builder strings.Builder
	for _, product := range products {
		builder.WriteString(fmt.Sprintf("%s\n", product))
	}
	return builder.String()
}

func (products Products) filter(predicate func(Product) bool) Products {
	var result Products
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

/* Write a generic filter "function" (not a method) that can be used to filter a collection of any type*/
func Filter[T comparable](list []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	/*
		var stationaryProducts = products.filter(func(p Product) bool {
			return p.Category == "Stationary"
		})
	*/
	var stationaryProducts = Filter(products, func(p Product) bool {
		return p.Category == "Stationary"
	})

	fmt.Println("Stationary Products")
	fmt.Println(stationaryProducts)

	nos := []int{3, 1, 4, 2, 5}
	var evenNos = Filter(nos, func(no int) bool {
		return no%2 == 0
	})
	fmt.Println(evenNos)
	// floats := []float32{4.5, 1.7, 4.8, 2.3, 5.4}

}
