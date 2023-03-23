package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genFib(stopCh)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	for fibNo := range ch {
		fmt.Println(fibNo)
	}
	fmt.Println("Done")
}

// producer
func genFib(stopCh chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
