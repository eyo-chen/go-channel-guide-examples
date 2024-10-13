package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	fmt.Println("Attempting to send to a closed channel...")
	ch <- 1
}
