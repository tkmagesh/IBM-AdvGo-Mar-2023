package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Something went wrong! :", err)
			return
		}
	}()

	fmt.Println("Async Execution")
	/*
		ch, errCh := divideAsync(100, 0)
		select {
		case result := <-ch:
			fmt.Println("Result :", result)
		case err := <-errCh:
			fmt.Println("Error :", err)
		}
	*/

	//making the error handling a choice
	ch, _ := divideAsync(100, 0)
	result := <-ch
	fmt.Println("Result :", result)

	/*
		fmt.Println("Sync Execution")
		result := divideSync(100, 0)
		fmt.Println("Result :", result)
	*/
}

func divideAsync(x, y int) (<-chan int, <-chan error) {
	ch := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				errCh <- err.(error)
				ch <- 0
			}
		}()
		time.Sleep(2 * time.Second)
		ch <- x / y
	}()
	return ch, errCh
}

func divideSync(x, y int) int {
	time.Sleep(2 * time.Second)
	return x / y
}
