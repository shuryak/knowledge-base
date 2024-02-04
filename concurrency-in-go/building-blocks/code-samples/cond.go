package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаём условие, используя стандартный sync.Mutex в качестве Locker
	c := sync.NewCond(&sync.Mutex{})
	// Создаём срез с нулевой длиной. Мы заранее знаем capacity — 10 элементов
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)

		// Мы снова входим в критическую секцию для условия, чтобы мы могли
		// изменить данные, относящиеся к условию
		c.L.Lock()
		// Симулируем удаление элемента из очереди путём изменения начала среза
		queue = queue[1:]

		fmt.Println("Removed from queue")

		// Выходим из критической секции условия, так как мы успешно удалили
		// элемент из очереди
		c.L.Unlock()

		// Сообщаем горутине, ждущей условие, что что-то произошло
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		// Входим в критическую секцию условия, вызывая Lock на Locker'е условия
		c.L.Lock()
		// Проверяем длину очереди в цикле. Это важно, поскольку сигнал не
		// означает, что произошло именно то, чего мы ждали, а только что ЧТО-ТО
		// произошло
		for len(queue) == 2 {
			// Вызов Wait, который приостановит (suspend) main горутину до тех
			// пор, пока не будет отправлен сигнал условия
			c.Wait()
		}

		fmt.Println("Adding to queue")

		queue = append(queue, struct{}{})
		// Создаём новую горутину, которая будет удалять элемент из очереди
		// через одну секунду
		go removeFromQueue(1 * time.Second)
		// Выходим из критической секции, так как мы успешно запушили в очередь
		c.L.Unlock()
	}
}
