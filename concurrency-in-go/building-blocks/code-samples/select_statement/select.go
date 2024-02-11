package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		// Закрываем канал по прошествии пяти секунд
		close(c)
	}()

	fmt.Println("Blocking on read...")

	select {
	// Пытаемся читать с канала. Следует обратить внимание, что для такого кода
	// не требуется оператор select, — мы могли бы просто написать <-c, но
	// напишем его для примера
	case <-c:
		fmt.Printf("Unblocking %v later.\n", time.Since(start))
	}
}
