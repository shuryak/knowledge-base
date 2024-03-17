package stage

import "sync"

func ToInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()

	return intStream
}

func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)

		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()

	return valueStream
}

func Take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
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

func TakeInt(done <-chan interface{}, valueStream <-chan int, num int) <-chan int {
	takeStream := make(chan int)

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

func PrimeFinder(done <-chan interface{}, valueStream <-chan int) <-chan int {
	primeStream := make(chan int)

	go func() {
		defer close(primeStream)

		for num := range valueStream {
			if slowIsPrime(num) {
				select {
				case <-done:
					return
				case primeStream <- num:
				}
			}
		}
	}()

	return primeStream
}

func FanInInt(done <-chan interface{}, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedStream := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()

		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

// FanIn - Здесь мы принимаем наш стандартный done-канал, чтобы разрешить отмену
// наших горутин, а затем принимаем произвольный слайс типа interface{} для
// объединения (fan-in)
func FanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	// Создаём sync.WaitGroup, чтобы мы могли дождаться, пока все каналы не
	// будут опустошены
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	// Создаём функцию multiplex, которая при передаче канала будет считывать из
	// канала и передавать считанное значение в канал multiplexedStream
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()

		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// Select из всех каналов.
	// Увеличиваем sync.WaitGroup на количество каналов, которые мы
	// мультиплексируем
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// Ожидание завершения всех чтений.
	// Создаём горутину для ожидания, пока все каналы, которые мы
	// мультиплексируем, не будут опустошены, чтобы мы могли закрыть канал
	// multiplexedStream
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func slowIsPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func fastIsPrime(number int) bool {
	if number <= 1 {
		return false
	}

	if number <= 3 {
		return true
	}

	if number%2 == 0 || number%3 == 0 {
		return false
	}

	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}

	return true
}
