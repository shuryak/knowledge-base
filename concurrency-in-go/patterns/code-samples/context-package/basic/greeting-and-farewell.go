package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	done := make(chan interface{})
	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(done); err != nil {
			fmt.Printf("%v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
		}
	}()

	wg.Wait()
}

func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}

	fmt.Printf("%s, мир!\n", greeting)

	return nil
}

func printFarewell(done <-chan interface{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}

	fmt.Printf("%s, мир!\n", farewell)

	return nil
}

func genGreeting(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "RU/RU":
		return "Привет", nil
	}

	return "", errors.New("unsupported locale")
}

func genFarewell(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "RU/RU":
		return "Пока", nil
	}

	return "", errors.New("unsupported locale")
}

func locale(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", errors.New("canceled")
	case <-time.After(1 * time.Minute):
	}

	return "RU/RU", nil
}
