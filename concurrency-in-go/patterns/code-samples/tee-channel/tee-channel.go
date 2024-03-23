package main

import "fmt"

func main() {
	tee := func(
		done <-chan interface{},
		in <-chan interface{},
	) (<-chan interface{}, <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			for val := range orDone(done, in) {
				// Нам нужны локальные копии out1 и out2, поэтому мы shadow'им
				// (затеняем) эти переменные
				out1, out2 := out1, out2

				// Мы собираемся использовать одно выражение select, чтобы
				// записи в out1 и out2 не блокировали друг друга. Чтобы
				// убедиться, что запись выполняется в оба, мы выполним две
				// итерации с select: по одной для каждого исходящего канала
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out1 <- val:
						// Как только мы выполнили запись в канал, мы
						// устанавливаем для его теневой копии значение nil,
						// чтобы дальнейшие записи были заблокированы и другой
						// канал мог продолжить работу
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()

		return out1, out2
	}

	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
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

func repeat(
	done <-chan interface{},
	values ...interface{},
) <-chan interface{} {
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

func take(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
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
