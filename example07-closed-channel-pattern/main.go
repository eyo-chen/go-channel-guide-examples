package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go sender(ch)
	go receiver(ch)

	time.Sleep(2 * time.Second)
}

func sender(ch chan<- int) {
	for i := 0; i < 2; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int) {
	// keep receiving data from the channel in an infinite loop
	for {
		value, ok := <-ch

		// if the channel is closed, break the loop
		if !ok {
			fmt.Println("Channel closed")
			break
		}

		fmt.Printf("Received value: %v\n", value)
	}
}

func RangeReceiver(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Received value: %v\n", value)
	}
}
