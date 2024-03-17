package main

import (
	"fan-out-fan-in/stage"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	randFn := func() interface{} {
		return rand.Intn(5_000_000_000)
	}

	done := make(chan interface{})
	defer close(done)

	randIntStream := stage.ToInt(done, stage.RepeatFn(done, randFn))

	numFinders := runtime.NumCPU()

	fmt.Printf("Spinning up %d prime finders.\n", numFinders)

	finders := make([]<-chan int, numFinders)

	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = stage.PrimeFinder(done, randIntStream)
	}

	start := time.Now()

	for prime := range stage.TakeInt(done, stage.FanInInt(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(start))
}
