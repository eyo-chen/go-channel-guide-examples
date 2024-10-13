package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		// opps, forgot to close the channel after sending all data
	}()

	// keep receiving data from the channel
	for v := range ch {
		fmt.Println(v)
	}
}
