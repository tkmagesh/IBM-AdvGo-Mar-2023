package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data2.dat")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	evenSum := 0
	oddSum := 0
	for scanner.Scan() {
		val := scanner.Text()
		if x, err := strconv.Atoi(val); err == nil {
			if x%2 == 0 {
				evenSum += x
			} else {
				oddSum += x
			}
		}
	}
	fmt.Println("Even Sum :", evenSum)
	fmt.Println("Odd Sum :", oddSum)
}
