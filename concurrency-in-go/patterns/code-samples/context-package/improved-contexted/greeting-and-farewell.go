package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)

			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()

	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("%s, мир!\n", greeting)

	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("%s, мир!\n", farewell)

	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "RU/RU":
		return "Привет", nil
	}

	return "", errors.New("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "RU/RU":
		return "Пока", nil
	}

	return "", errors.New("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	// Здесь мы проверяем, установлен ли в нашем `Context` дедлайн. Если
	// установлен, и наши системные часы превысили дедлайн, мы просто
	// возвращаемся из функции со специальной ошибкой DeadlineExceeded,
	// определённой в пакете context
	if deadline, ok := ctx.Deadline(); ok {
		if time.Now().Add(1 * time.Minute).After(deadline) {
			return "", context.DeadlineExceeded
		}
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}

	return "RU/RU", nil
}
