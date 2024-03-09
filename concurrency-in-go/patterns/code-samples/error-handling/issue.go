package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)

		go func() {
			defer close(responses)

			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					// Здесь мы видим, что горутина делает всё возможное, чтобы
					// сигнализировать о наличии ошибки. Что ещё она может
					// сделать? Она не может передать эту ошибку обратно.
					// Сколько ошибок считать слишком большим количеством?
					// Продолжает ли она после сигнала об ошибке отправлять
					// запросы?
					fmt.Println(err)
					continue
				}

				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()

		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://google.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
