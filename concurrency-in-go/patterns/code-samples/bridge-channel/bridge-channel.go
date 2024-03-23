package main

import "fmt"

func main() {
	bridge := func(
		done <-chan interface{},
		chanStream <-chan <-chan interface{},
	) <-chan interface{} {
		// Это канал, который вернёт все значения из bridge
		valStream := make(chan interface{})

		go func() {
			defer close(valStream)

			// Этот цикл отвечает за извлечение каналов из chanStream и
			// предоставление их для использования во вложенном цикле
			for {
				var stream <-chan interface{}

				select {
				case maybeStream, ok := <-chanStream:
					if !ok {
						return
					}

					stream = maybeStream
				case <-done:
					return
				}

				// Этот вложенный цикл отвечает за считывание значений с
				// заданного канала и повторение этих значений в valStream.
				// Когда поток, который мы в данный момент перебираем,
				// закрывается, мы выходим из этого вложенного цикла, выполняя
				// считывание из этого канала, и переходим к следующей итерации
				// цикла, выбирающего каналы для чтения. Это обеспечивает
				// непрерывный поток значений
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()

		return valStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))

		go func() {
			defer close(chanStream)

			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()

		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}

	fmt.Println()
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)

		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}

				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()

	return valStream
}
