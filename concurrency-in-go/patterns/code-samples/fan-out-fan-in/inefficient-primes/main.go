package main

import (
	"fan-out-fan-in/stage"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randFn := func() interface{} {
		return rand.Intn(5_000_000_000)
	}

	done := make(chan interface{})
	defer close(done)

	randIntStream := stage.ToInt(done, stage.RepeatFn(done, randFn))

	fmt.Println("Primes:")

	start := time.Now()

	for prime := range stage.TakeInt(done, stage.PrimeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(start))
}
