package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	// fmt.Println(divide(100, 7))
	divisor := 0
	q, _, err := divide(100, divisor)
	if err == DivideByZeroError {
		fmt.Println("Please donot attempt to divide by zero!")
		return
	}
	if err != nil {
		fmt.Println("something went wrong :", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d \n", divisor, q)
}

/*
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}
