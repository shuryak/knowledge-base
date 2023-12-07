package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	// Каждое направление представляется названием (left/right) и числом людей,
	// пытающихся переместиться в этом направлении (dir)
	type direction struct {
		name  string
		value *int32
	}

	// tryDirection позволяет человеку совершить попытку переместиться в
	// заданном направлении и возвращает булево значение успешности этого
	// перемещения.
	tryDirection := func(d direction, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %s", d.name)
		// Во-первых, мы заявляем о нашем намерении переместиться в данном
		// направлении увеличением значения этого направления на 1.
		// Всё, что сейчас необходимо знать про atomic — то, что операции в этом
		// пакете атомарные.
		atomic.AddInt32(d.value, 1)
		// Чтобы продемонстрировать livelock, каждый человек должен двигаться с
		// одинаковым показателем скорости (ритмом).
		// takeStep() симулирует постоянный ритм для всех сторон.
		takeStep()
		if atomic.LoadInt32(d.value) == 1 {
			fmt.Fprint(out, ". Success!\n")
			return true
		}
		takeStep()
		// Здесь человек понимает, что он не сможет пройти и сдаётся. Мы
		// показываем это уменьшением значения этого направления на 1.
		atomic.AddInt32(d.value, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDirection(direction{"left", &left}, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDirection(direction{"right", &right}, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			fmt.Printf(out.String())
		}()
		defer walking.Done()
		fmt.Fprintf(&out, "%s is trying to scoot:", name)
		// Искусственное ограничение на число попыток для того, чтобы программа
		// в конечном итоге завершилась. В программе с livelock'ом такого
		// ограничения может и не быть, вот почему это проблема!
		for i := 0; i < 5; i++ {
			// Сначала человек пытается пройти влево, и если это движение
			// завершается неудачей, он пытается пройти вправо.
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%s tosses her hands up in exasperation!\n", name)
	}

	// Эта переменная предоставляет программе способ подождать до тех пор, пока
	// оба человека не смогут пройти или сдаться.
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
