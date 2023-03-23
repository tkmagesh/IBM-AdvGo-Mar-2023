package main

import (
	"fmt"
	"time"
)

// share memory by communicating

/*
func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()

	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(4 * time.Second)
	ch <- x + y
	wg.Done()
}
*/

// the above without waitgroup
func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch //blocked
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	time.Sleep(4 * time.Second)
	ch <- x + y //non-blocking operation
}
