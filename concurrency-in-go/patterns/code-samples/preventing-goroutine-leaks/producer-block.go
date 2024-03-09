package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)

		go func() {
			// Выводим сообщение, когда горутина успешно завершается
			defer fmt.Println("newRandStream closure exited.")

			defer close(randStream)

			for {
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()

	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i+1, <-randStream)
	}
}
