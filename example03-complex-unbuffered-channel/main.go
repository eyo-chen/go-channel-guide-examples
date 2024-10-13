package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go receiver(ch)
	fmt.Println("Sending value 1 to channel")
	ch <- 1

	fmt.Println("Sending value 2 to channel")
	ch <- 2

	fmt.Println("Sending value 3 to channel")
	ch <- 3

	// sleep 3 seconds to ensure the goroutine has time to finish goroutine
	time.Sleep(3 * time.Second)
}

func receiver(ch chan int) {
	val := 0
	for val != 3 {
		val = <-ch

		// sleep 5 seconds if the received value is 1
		if val == 1 {
			time.Sleep(5 * time.Second)
		}

		fmt.Println("Received:", val)
	}
}
