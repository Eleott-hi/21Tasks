package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// multiplex function accepts a variable number of input channels and returns a single output channel.
func multiplex(channels ...chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan interface{}) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// Sample data: 3 input channels
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	// Test the multiplex function
	out := multiplex(ch1, ch2, ch3)

	// Generate and send random values to input channels
	go func() {
		for i := 0; i < 10; i++ {
			val := rand.Intn(100)
			select {
			case ch1 <- 1000 + val:
			case ch2 <- 2000 + val:
			case ch3 <- 3000 + val:
			}
			time.Sleep(time.Millisecond * 100) // Simulate random intervals
		}
		close(ch1)
		close(ch2)
		close(ch3)
	}()

	// Receive and print values from the output channel
	for val := range out {
		fmt.Println(val)
	}
}
