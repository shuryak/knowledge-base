package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int

	select {
	// Этот оператор case никогда не будет разблокирован, поскольку он хочет
	// читать с nil-канала
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}
