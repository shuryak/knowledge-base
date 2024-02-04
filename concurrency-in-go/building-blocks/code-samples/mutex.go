package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()

		count++

		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		// Запрашиваем исключительное использование критической секции — в этом
		// случае переменная count, защищённая Mutex'ом, блокируется
		lock.Lock()
		// Указываем, что мы закончили с блокировкой критической секции
		defer lock.Unlock()

		count--

		fmt.Printf("Decrementing: %d\n", count)
	}

	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			decrement()
		}()
	}

	wg.Wait()

	fmt.Println("Arithmetic complete.")
}
