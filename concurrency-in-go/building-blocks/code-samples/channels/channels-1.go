package main

import "fmt"

func main() {
	stringStream := make(chan string)

	go func() {
		stringStream <- "Hello channels!" // Пишем строчку в канал stringStream
	}()

	// Читаем строчку из канала stringStream и выводим её в stdout
	fmt.Println(<-stringStream)
}
