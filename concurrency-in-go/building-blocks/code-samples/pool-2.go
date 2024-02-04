package main

import (
	"fmt"
	"sync"
)

func main() {
	calculatorsNum := 0

	calcPool := &sync.Pool{
		New: func() interface{} {
			calculatorsNum++
			mem := make([]byte, 1024)
			return &mem // мы храним адрес слайса байтов
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const workersNum = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(workersNum)

	for i := 0; i < workersNum; i++ {
		go func() {
			defer wg.Done()

			// Утверждаем, что тип — это указатель на слайс байтов
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			// Предположим, что что-то интересное, но быстрое делается с этой
			// памятью
		}()
	}

	wg.Wait()

	fmt.Printf("%d calculators were created\n", calculatorsNum)
}
