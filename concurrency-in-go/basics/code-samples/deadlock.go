package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		// Попытка войти в критическую секцию для входящего значения
		v1.mu.Lock()
		// Выход из критической секции, когда произойдёт return
		defer v1.mu.Unlock()

		// Спим для симуляции работы (это спровоцирует deadlock)
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%d\n", v1.value+v2.value)

	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
