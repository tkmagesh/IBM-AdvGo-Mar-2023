package main

import "fmt"

func sumInts(nos ...int) int {
	var result int
	for _, no := range nos {
		result += no
	}
	return result
}

func sumFloats(nos ...float32) float32 {
	var result float32
	for _, no := range nos {
		result += no
	}
	return result
}

/*
func sum[T int | float32](nos ...T) T {
	var result T
	for _, no := range nos {
		result += no
	}
	return result
}
*/

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

func sum[T Numeric](nos ...T) T {
	var result T
	for _, no := range nos {
		result += no
	}
	return result
}

func main() {
	ints := []int{4, 1, 4, 2, 5}
	// fmt.Println(sumInts(ints...))
	fmt.Println(sum(ints...))
	floats := []float32{4.5, 1.7, 4.8, 2.3, 5.4}
	// fmt.Println(sumFloats(floats...))
	fmt.Println(sum(floats...))
}
