package main

import (
	"fmt"
	"sync"
)

func main() {
	// Определяем тип Button, который содержит Cond — Clicked
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// Определяем функцию, которая позволит регистрировать функции для обработки
	// сигналов из Cond. Каждый обработчик запускается в своей горутине, и
	// subscribe не завершится, пока не будет подтверждено, что эта горутина
	// запущена
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		// Устанавливаем обработчик для того, когда кнопка мыши отпущена. Он, в
		// свою очередь, вызывает Broadcast для Clicked Cond, чтобы сообщить
		// всем обработчикам, что кнопка мыши была нажата (более надёжная
		// реализация сначала проверила бы, что она была нажата)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	// Создаём WaitGroup только для того, чтобы быть уверенным, что программа не
	// завершится до полного вывода в stdout
	var clickRegistered sync.WaitGroup

	clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	// Симулируем то, как пользователь отпускает кнопку мыши после того, как
	// нажал на кнопку в GUI
	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
