package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var c1, c2 <-chan int

	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("Control in default clause after %v\n", time.Since(start))
	}
}
