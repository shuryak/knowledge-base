package main

import (
	"fmt"
	"sync"
)

func main() {
	// Переменная, которая будет разрешать/запрещать доступ к памяти переменной
	// data
	var memoryAccess sync.Mutex
	var data int
	go func() {
		// Объявляем, что пока мы не объявим обратное, горутина должна иметь
		// исключительный доступ к этой памяти
		memoryAccess.Lock()
		data++
		// Объявляем, что горутина завершила работу с этой памятью
		memoryAccess.Unlock()
	}()
	// Снова объявляем, что следующие условные выражения должны иметь
	// исключительный доступ к памяти переменной data
	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is %d.\n", data)
	} else {
		fmt.Printf("the value is %d.\n", data)
	}
	// Снова объявляем, что работа с памятью переменной data завершена
	memoryAccess.Unlock()
}
