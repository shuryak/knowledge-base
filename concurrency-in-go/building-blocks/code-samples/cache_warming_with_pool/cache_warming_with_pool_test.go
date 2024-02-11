package main

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

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{New: connectToService}

	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return p
}

// Функция, которая симулирует создание подключения к сервису.
func connectToService() interface{} {
	time.Sleep(time.Second)

	return struct{}{}
}

// Вариант с sync.Pool
func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen, %v", err)
		}

		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			svcConn := connPool.Get()
			_, _ = fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			_ = conn.Close()
		}
	}()

	return &wg
}
