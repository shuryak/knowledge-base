package queuing

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
)

func BenchmarkUnbufferedWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}

func BenchmarkBufferedWrite(b *testing.B) {
	bufferedFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferedFile))
}

func tmpFileOrFatal() *os.File {
	file, err := os.CreateTemp("", "tmp")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return file
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()

	for bt := range take(done, repeat(done, byte(0)), b.N) {
		_, err := writer.Write([]byte{bt.(byte)})
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
}

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)

		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()

	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}
