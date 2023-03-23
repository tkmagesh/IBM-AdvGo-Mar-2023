package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	ch3 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("ch3 : ", <-ch3)
	}()

	/*
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			fmt.Println(<-ch1)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			fmt.Println(<-ch2)
			wg.Done()
		}()
		wg.Wait()
	*/
	for i := 0; i < 3; i++ {
		select {
		case data1 := <-ch1:
			fmt.Println(data1)
		case data2 := <-ch2:
			fmt.Println(data2)
		case ch3 <- 300:
			fmt.Println("Attempt to send data to ch3 succeeded!")
		}
	}
}
