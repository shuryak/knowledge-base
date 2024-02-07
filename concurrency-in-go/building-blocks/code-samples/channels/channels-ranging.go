package main

import "fmt"

func main() {
	intStream := make(chan int)

	go func() {
		// Обеспечиваем, что канал будет закрыт перед выходом из горутины. Очень
		// распространённый паттерн
		defer close(intStream)

		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	// Перебор канала intStream
	for i := range intStream {
		fmt.Printf("%d ", i)
	}

	fmt.Println()
}
