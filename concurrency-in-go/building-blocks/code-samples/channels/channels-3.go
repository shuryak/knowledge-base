package main

import "fmt"

func main() {
	stringStream := make(chan string)

	go func() {
		// Обеспечиваем невозможность получения значения из канала stringStream
		// из main-горутины
		if 0 != 1 {
			return
		}

		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}
