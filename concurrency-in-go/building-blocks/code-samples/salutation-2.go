package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) { // принимаем параметр как в обычную функцию
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
		// передаём текущую итерационную переменную в замыкание, создаётся копия
		// структуры строки
	}
	wg.Wait()
}
