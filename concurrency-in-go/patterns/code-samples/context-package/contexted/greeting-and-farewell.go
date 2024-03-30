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

	// Здесь main создаёт новый Context через context.Background и заворачивает
	// его в context.WithCancel, чтобы предоставить возможность отмен
	// (cancellations)
	ctx, cancel := context.WithCancel(context.Background()) // (1)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)

			// Эта строчка спровоцирует то, что main отменит Context, если из
			// printGreeting вернётся ошибка
			cancel() // (2)
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
	// Здесь genGreeting заворачивает свой Context через context.WithTimeout.
	// Это автоматически отменит возвращаемый Context через 1 секунду, тем самым
	// отменив все дочерние элементы, куда этот возвращаемый Context будет
	// передан, а именно locale
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // (3)
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
	select {
	case <-ctx.Done():
		// Эта строчка возвращает причину, по которой контекст был отменён. Эта
		// ошибка будет поочерёдно всплывать вплоть до main, что приведёт к
		// отмене в (2)
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}

	return "RU/RU", nil
}
