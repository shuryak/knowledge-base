package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %d\n", id)
	}

	const greetersCount = 5

	var wg sync.WaitGroup

	wg.Add(greetersCount)

	for i := 0; i < greetersCount; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
}
