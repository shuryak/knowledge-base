package main

import "fmt"

func main() {
	c1 := make(chan interface{})
	close(c1)

	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int

	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}
