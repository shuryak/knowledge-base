package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	// нам нужна горутина, которая никогда не завершится, чтобы хранить их в
	// определённом количестве в памяти для измерений
	noop := func() {
		wg.Done()
		<-c
	}

	// задаём число горутин для создания.
	// будем использовать закон больших чисел, чтобы асимптотически приблизиться
	// к размеру горутины
	const goroutinesCount = 1e4
	wg.Add(goroutinesCount)
	// измеряем объём потребляемой памяти перед созданием горутин
	before := memConsumed()
	for i := 0; i < goroutinesCount; i++ {
		go noop()
	}
	wg.Wait()
	// измеряем объём потребляемой памяти после создания горутин
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/goroutinesCount/1000)
}
