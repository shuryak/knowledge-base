package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	// Второй параметр функции имеет тип sync.Locker. Это интерфейс, который
	// имеет два метода: Lock и Unlock. Mutex и RWMutex ему удовлетворяют
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			l.Lock()
			l.Unlock()
			// producer спит 1 секунду, что делает его менее активным, чем горутины observer
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()

		l.Lock()
		defer l.Unlock()
	}

	test := func(observersCount int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup

		wg.Add(observersCount + 1) // observers + 1 producer

		beginTestTime := time.Now()

		go producer(&wg, mutex)

		for i := 0; i < observersCount; i++ {
			go observer(&wg, rwMutex)
		}

		wg.Wait()

		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex

	_, _ = fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		_, _ = fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}
