package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() { // отправитель
		defer wg.Done()
		// ждём пока не будет команды начинать.
		// мы не хотим, чтобы затраты на настройку и запуск каждой горутины
		// учитывались при измерении
		<-begin
		for i := 0; i < b.N; i++ {
			// отправляем сообщения в горутину-получатель.
			// struct{}{} называется пустой структурой и не занимает памяти
			c <- token
		}
	}

	receiver := func() { // получатель
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c // получаем сообщение и ничего с ним не делаем
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // запускаем таймер производительности
	close(begin)   // даём команду начинать двум горутинам
	wg.Wait()
}
