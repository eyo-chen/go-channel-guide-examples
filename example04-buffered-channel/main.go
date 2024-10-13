package main

import (
	"fmt"
	"time"
)

func main() {
	// create a buffered channel with capacity 2
	ch := make(chan int, 2)

	// start timing to calculate the time elapsed
	start := time.Now()

	go sender(ch, start)
	go receiver(ch, start)

	// delay for two goroutines to complete
	time.Sleep(6 * time.Second)
}

func sender(ch chan<- int, start time.Time) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Sending %d at %s\n", i, time.Since(start))
		ch <- i
		fmt.Printf("Sent %d at %s\n", i, time.Since(start))
	}
}

func receiver(ch <-chan int, start time.Time) {
	// delay start to demonstrate buffer filling up
	time.Sleep(2 * time.Second)

	for i := 1; i <= 3; i++ {
		val := <-ch
		fmt.Printf("Received %d at %s\n", val, time.Since(start))

		// sleep to simulate slow receiver
		time.Sleep(1000 * time.Millisecond)
	}
}
