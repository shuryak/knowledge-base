package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0

loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// Симулируем работу
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %d cycles of work before signalled to stop.\n", workCounter)
}
