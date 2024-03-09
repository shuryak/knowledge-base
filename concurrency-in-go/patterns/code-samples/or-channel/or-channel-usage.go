package main

import (
	"fmt"
	"time"
)

func main() {
	// Эта функция просто создаёт канал, который закрывается по прошествии
	// времени, указанного в after
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	// Здесь мы запоминаем приблизительное время, когда канал из функции or
	// начинает блокировать main-горутину
	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	// Выводим время, которое потребовалось для выполнения считывания из канала,
	// возвращённого функцией or
	fmt.Printf("done after %v\n", time.Since(start))
}
