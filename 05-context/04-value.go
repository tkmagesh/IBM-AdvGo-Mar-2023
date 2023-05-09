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
	valCtx := context.WithValue(rootCtx, "root-key", "root-val")
	cancelCtx, cancel := context.WithCancel(valCtx) // context used to send the cancel signal
	defer cancel()
	go func() {
		fmt.Println("Hit ENTER to stop")
		fmt.Scanln()
		cancel()
	}()

	fmt.Println("[main] starting the goroutine")
	wg.Add(1) // increment the wg counter by 1
	go fn(cancelCtx, wg)
	wg.Wait() // block until the wg counter becomes 0
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	counter := 0

	fnValCtx := context.WithValue(ctx, "fn-key", "fn-val")

	timeoutCtx1, cancel := context.WithTimeout(fnValCtx, 5*time.Second)
	defer cancel()

	fmt.Println("[fn] root-key : ", ctx.Value("root-key"))
	fmt.Println("[fn] starting f1")
	wg.Add(1)
	go f1(timeoutCtx1, wg)

	timeoutCtx2, cancel := context.WithTimeout(fnValCtx, 7*time.Second)
	defer cancel()

	fmt.Println("[fn] starting f2")
	wg.Add(1)
	go f2(timeoutCtx2, wg)

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

func f1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	counter := 0
	fmt.Println("[f1] root-key : ", ctx.Value("root-key"))
	fmt.Println("[f1] fn-key : ", ctx.Value("fn-key"))
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[f1] Cancel signal received... exiting the goroutine")
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			counter += 2
			fmt.Println("[f1] : ", counter)
		}

	}
}

func f2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	counter := 0
	fmt.Println("[f2] root-key : ", ctx.Value("root-key"))
	fmt.Println("[f2] fn-key : ", ctx.Value("fn-key"))
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[f2] Cancel signal received... exiting the goroutine")
			break LOOP
		default:
			time.Sleep(700 * time.Millisecond)
			counter += 7
			fmt.Println("[f2] : ", counter)
		}

	}
}
