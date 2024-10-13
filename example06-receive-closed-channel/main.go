package main

import "fmt"

func main() {
	ch := make(chan int)

	close(ch)

	fmt.Println("Reading from a closed channel:")
	value, ok := <-ch
	fmt.Printf("Value: %v, Channel open: %v\n", value, ok)
}
