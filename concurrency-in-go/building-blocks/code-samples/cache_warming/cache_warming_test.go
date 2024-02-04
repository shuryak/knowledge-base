package cache_warming

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}

		if _, err := io.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}

		_ = conn.Close()
	}
}

// Функция, которая симулирует создание подключения к сервису.
func connectToService() interface{} {
	time.Sleep(time.Second)

	return struct{}{}
}

// Посмотрим, насколько эффективным будет сетевой сервис, если для каждого запроса мы будем открывать новое соединение с
// сервисом. Напишем сетевой обработчик, который открывает соединение с другим сервисом для каждого соединения, принимаемого
// сетевым обработчиком. Чтобы упростить бенчмаркинг, мы разрешим только одно одновременное соединение
func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}

		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			connectToService()
			_, _ = fmt.Fprintln(conn, "")
			_ = conn.Close()
		}
	}()

	return &wg
}
