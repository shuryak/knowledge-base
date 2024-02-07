package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			<-begin // Здесь горутина ждёт, пока ей сообщат продолжать

			fmt.Printf("%d has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")

	// Закрываем канал, тем самым разблокируя все горутины одновременно
	close(begin)
	wg.Wait()
}
