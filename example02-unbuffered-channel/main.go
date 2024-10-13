package main

import (
	"fmt"
	"time"
)

func main() {
	// create an unbuffered channel
	ch := make(chan int)

	go func() {
		fmt.Println("Sending value 1 to channel")
		ch <- 1
		fmt.Println("After sending value 1")
	}()

	// sleep for 3 seconds to ensure the goroutine has time to send the value
	time.Sleep(3 * time.Second)

	fmt.Println("Receiving value from channel")
	val := <-ch
	fmt.Println(val)

	// sleep 1 second to ensure goroutine has time to finish goroutine
	time.Sleep(1 * time.Second)
}
