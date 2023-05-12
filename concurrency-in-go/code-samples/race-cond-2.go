package main

import "fmt"

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %d.\n", data)
	}
}
