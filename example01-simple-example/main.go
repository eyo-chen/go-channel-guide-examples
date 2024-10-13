package main

import "fmt"

func main() {
	// create a channel
	ch := make(chan int)

	// start a goroutine to send data to the channel
	go func() {
		// send data to the channel
		ch <- 1
	}()

	// receive data from the channel
	val := <-ch
	fmt.Println(val)
}
