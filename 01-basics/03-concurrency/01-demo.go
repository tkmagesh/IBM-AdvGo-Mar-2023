package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")

	wg.Add(1)
	go fn(wg)

	wg.Add(1)
	go f1(wg)

	wg.Wait()
	fmt.Println("main completed")
}

func fn(wg *sync.WaitGroup) {
	fmt.Println("fn invoked")
	wg.Done()
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(10 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}
