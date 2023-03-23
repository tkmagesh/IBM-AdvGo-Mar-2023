package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = "This is a string"
	x = true
	x = 99.88
	fmt.Println(x)
}
