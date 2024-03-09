package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Создаём тип, который включает в себя как *http.Response, так и ошибку,
	// возможную в цикле нашей горутины
	type Result struct {
		Error    error
		Response *http.Response
	}

	// Эта функция возвращает канал, из которого можно считывать результаты
	// итераций нашего цикла
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				resp, err := http.Get(url)

				// Создаём инстанс Result с необходимым набором полей
				result := Result{
					Error:    err,
					Response: resp,
				}

				select {
				case <-done:
					return
				// Пишем result в канал
				case results <- result:
				}
			}
		}()

		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		// Здесь, в нашей main-горутине, мы можем разумнее обрабатывать ошибки,
		// возникающие в горутине, запущенной из checkStatus, и в рамках более
		// обширного контекста
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			continue
		}

		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
