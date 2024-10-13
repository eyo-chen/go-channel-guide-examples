package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type Channel[M any] struct {
	queue    *list.List
	capacity int
	cond     *sync.Cond
	closed   bool
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		queue:    list.New(),
		capacity: capacity,
		cond:     sync.NewCond(&sync.Mutex{}),
		closed:   false,
	}
}

func (c *Channel[M]) Send(message M) error {
	// lock the channel
	c.cond.L.Lock()

	// unlock the channel after the function returns
	defer c.cond.L.Unlock()

	// if the channel is closed, return an error
	if c.closed {
		return errors.New("send on closed channel")
	}

	// if the channel is full, wait for the channel to have space
	for c.queue.Len() == c.capacity {
		// Wait() will:
		// (1) unlock the channel
		// (2) sleep until the channel is signaled
		c.cond.Wait()
	}

	// add the message to the queue
	c.queue.PushBack(message)

	// broadcast to all the goroutines waiting on the channel
	c.cond.Broadcast()
	return nil
}

func (c *Channel[M]) Receive() (M, bool) {
	// lock the channel
	c.cond.L.Lock()

	// unlock the channel after the function returns
	defer c.cond.L.Unlock()

	// if the channel is closed, return a zero value and false
	if c.closed {
		var zero M
		return zero, false
	}

	// increment the capacity and broadcast to all the goroutines waiting on the channel
	// so that the sender can send a new message
	c.capacity++
	c.cond.Broadcast()

	// if the queue is empty, wait for the channel to have a message
	for c.queue.Len() == 0 {
		// Wait() will:
		// (1) unlock the channel
		// (2) sleep until the channel is signaled
		c.cond.Wait()
	}

	// remove the message from the queue
	message := c.queue.Remove(c.queue.Front()).(M)

	// decrement the capacity
	c.capacity--

	// broadcast to all the goroutines waiting on the channel
	c.cond.Broadcast()
	return message, true
}

func (c *Channel[M]) Close() error {
	// lock the channel
	c.cond.L.Lock()

	// unlock the channel after the function returns
	defer c.cond.L.Unlock()

	// if the channel is already closed, return an error
	if c.closed {
		return errors.New("close on closed channel")
	}

	// set the channel to closed
	c.closed = true

	// broadcast to all the goroutines waiting on the channel
	c.cond.Broadcast()
	return nil
}

func main() {
	ch := NewChannel[int](0)

	go func() {
		for i := 1; i <= 5; i++ {
			err := ch.Send(i)
			if err != nil {
				fmt.Printf("Send error: %v\n", err)
				return
			}
			fmt.Printf("Sent: %d\n", i)
		}
		ch.Close()
	}()

	for {
		value, ok := ch.Receive()
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Printf("Received: %d\n", value)
	}
}
