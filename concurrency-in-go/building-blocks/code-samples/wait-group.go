package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Вызываем Add с аргументом 1, чтобы указать, что одна горутина будет
	// запущена
	wg.Add(1)
	go func() {
		// Вызываем Done используя defer, чтобы убедиться в том, что перед
		// выходом из замыкания горутины, мы сообщим WaitGroup'е, что горутина
		// завершила работу
		defer wg.Done()
		fmt.Println("First goroutine sleeping...")
		time.Sleep(time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Second goroutine sleeping...")
		time.Sleep(2 * time.Second)
	}()

	// Вызываем Wait, чтобы блокировать основную (main) горутину до тех пор,
	// пока все горутины не сообщат, что они закончили работу
	wg.Wait()

	fmt.Println("All goroutines complete.")
}
