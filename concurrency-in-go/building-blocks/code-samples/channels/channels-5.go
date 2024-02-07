package main

import "fmt"

func main() {
	intStream := make(chan int)
	close(intStream)

	integer, ok := <-intStream // читаем из закрытого канала

	fmt.Printf("(%v): %d\n", ok, integer)
}
