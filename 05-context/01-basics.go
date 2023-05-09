package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx) // context used to send the cancel signal
	wg.Add(1)                                        // increment the wg counter by 1
	fmt.Println("[main] starting the goroutine")
	go f1(cancelCtx, wg)
	time.Sleep(5 * time.Second)
	fmt.Println("[main] timeout occured. sending the cancel signal")
	cancel()
	wg.Wait() // block until the wg counter becomes 0
}

func f1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	counter := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[fn] Cancel signal received... exiting the goroutine")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			counter++
			fmt.Println("[fn] : ", counter)
		}

	}
}
