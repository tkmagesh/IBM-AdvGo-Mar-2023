package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenSumCh := make(chan int)
	oddSumCh := make(chan int)

	sourceWg := &sync.WaitGroup{}
	sourceWg.Add(1)
	go source("data1.dat", dataCh, sourceWg)
	sourceWg.Add(1)
	go source("data2.dat", dataCh, sourceWg)

	processWg := &sync.WaitGroup{}
	processWg.Add(1)
	go splitter(dataCh, evenCh, oddCh, processWg)
	processWg.Add(1)
	go sum(evenCh, evenSumCh, processWg)
	processWg.Add(1)
	go sum(oddCh, oddSumCh, processWg)
	processWg.Add(1)
	go merger("result.txt", evenSumCh, oddSumCh, processWg)

	sourceWg.Wait()
	close(dataCh)

	processWg.Wait()
	fmt.Println("Done")
}

func source(fileName string, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		if no, err := strconv.Atoi(val); err == nil {
			dataCh <- no
		}
	}
}

func splitter(dataCh, evenCh, oddCh chan int, wg *sync.WaitGroup) {
	defer close(evenCh)
	defer close(oddCh)
	defer wg.Done()
	for no := range dataCh {
		if no%2 == 0 {
			evenCh <- no
		} else {
			oddCh <- no
		}
	}
}

func sum(ch chan int, resultCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for no := range ch {
		result += no
	}
	resultCh <- result
}

func merger(fileName string, evenSumCh, oddSumCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for i := 0; i < 2; i++ {
		select {
		case evenSum := <-evenSumCh:
			fmt.Fprintf(file, "Even Sum : %d\n", evenSum)
		case oddSum := <-oddSumCh:
			fmt.Fprintf(file, "Odd Sum : %d\n", oddSum)
		}
	}
}
