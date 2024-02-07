package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// Создаём буфер в памяти, чтобы ослабить недетерминированный характер
	// вывода. Это не даёт нам никаких гарантий, однако это немного быстрее, чем
	// писать в stdout напрямую
	var stdoutBuff bytes.Buffer

	// Этим мы гарантируем, что содержимое буфера будет записано в stdout перед
	// завершением процесса
	defer stdoutBuff.WriteTo(os.Stdout)

	// Создаём буферизированный канал с вместимостью четыре
	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d.\n", i)
			intStream <- i
		}
	}()

	for i := range intStream {
		_, _ = fmt.Fprintf(&stdoutBuff, "Received %d.\n", i)
	}
}
