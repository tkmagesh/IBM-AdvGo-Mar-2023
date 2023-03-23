package main

import (
	"fmt"
	"time"
)

// Solve this without using channels ( but using pointer & waitgroups )

/*
// consumer
func main() {
	ch := make(chan int)
	go genFib(ch)
	for fibNo := range ch {
		fmt.Println(fibNo)
	}
	fmt.Println("Done")
}

// producer
func genFib(ch chan int) {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		ch <- x
		time.Sleep(500 * time.Millisecond)
		x, y = y, x+y
	}
	close(ch)
}
*/

// consumer
func main() {
	ch := genFib()
	/*
		for fibNo := range ch {
			fmt.Println(fibNo)
		}
	*/
	for {
		if fibNo, isOpen := <-ch; isOpen {
			fmt.Println(fibNo)
			continue
		}
		break
	}
	fmt.Println("Done")
}

// producer
func genFib() chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < 10; i++ {
			ch <- x
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
		close(ch)
	}()
	return ch
}
