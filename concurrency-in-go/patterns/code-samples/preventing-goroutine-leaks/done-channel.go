package main

import (
	"fmt"
	"time"
)

func main() {
	// В функции doWork ожидаем канал done. По соглашению он должен быть первым
	// параметром
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					// Делаем что-то интересное
					fmt.Println(s)
				// Здесь мы видим повсеместно используемый паттерн for-select.
				// Одним из наших условий является проверка того, был ли сигнал
				// горутине из канала done. Если да, то мы возвращаемся из
				// горутины
				case <-done:
					return
				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	// Здесь создаём другую горутину, которая отменит горутину, созданную в
	// doWork, если пройдёт более одной секунды
	go func() {
		// Останавливаем операцию через 1 секунду
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine.")
		close(done)
	}()

	// Точка соединения (join point) горутины, созданной в doWork с
	// main-горутиной
	<-terminated

	fmt.Println("Done.")
}
