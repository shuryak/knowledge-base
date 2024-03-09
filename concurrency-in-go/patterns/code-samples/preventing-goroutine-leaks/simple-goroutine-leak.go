package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)

			for s := range strings {
				// Делаем что-то интересное
				fmt.Println(s)
			}
		}()

		return completed
	}

	doWork(nil)

	// Здесь может быть проделана ещё какая-либо работа

	fmt.Println("Done.")
}
