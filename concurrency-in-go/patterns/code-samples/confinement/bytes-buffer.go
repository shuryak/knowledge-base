package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			_, _ = fmt.Fprintf(&buff, "%c", b)
		}

		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	data := []byte("golang")

	go printData(&wg, data[:3]) // передаём слайс из первых трёх байт data
	go printData(&wg, data[3:]) // передаём слайс из последних трёх байт data

	wg.Wait()
}
