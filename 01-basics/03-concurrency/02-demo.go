package main

import (
	"fmt"
	"sync"
	"time"
)

// Communicating By Sharing Memory

/* var result int

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()

	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup) {
	time.Sleep(4 * time.Second)
	result = x + y
	wg.Done()
} */

func main() {
	result := 0
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go add(100, 200, wg, &result)
	wg.Wait()

	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, result *int) {
	time.Sleep(4 * time.Second)
	*result = x + y
	wg.Done()
}
