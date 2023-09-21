package tcp

import (
	"context"
	"testing"
	"time"

	"github.com/batazor/word_of_wisdow/internal/pkg/logger"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestPing(t *testing.T) {
	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new logger
	log, err := logger.New()
	if err != nil {
		t.Fatal(err)
	}

	// Create a new TCP server
	server, err := NewServer(ctx, "localhost:9090", log)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new TCP client
	client, err := NewClient("localhost:9090")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Send `ping` to the server
	_, err = client.Conn.Write([]byte("ping"))
	if err != nil {
		t.Error(err)
		return
	}

	// Read data from the server
	select {
	case v := <-server.ReadCh:
		if string(v) != "ping" {
			t.Fatal("expected ping, got", string(v))
		}
	case <-time.After(3 * time.Second):
		t.Fatal("timeout")
	}
}

func TestPong(t *testing.T) {
	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new logger
	log, err := logger.New()
	if err != nil {
		t.Fatal(err)
	}

	// Create a new TCP server
	server, err := NewServer(ctx, "localhost:9090", log)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new TCP client
	client, err := NewClient("localhost:9090")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Send `pong` to the client
	go func() {
		time.Sleep(100 * time.Millisecond)
		err = server.Send([]byte("pong\n"))
		if err != nil {
			t.Error(err)
			return
		}
	}()

	// Read data from the client
	select {
	case v := <-client.ReadCh:
		if string(v) != "pong\n" {
			t.Fatal("expected pong, got", string(v))
		}
	case <-time.After(3 * time.Second):
		t.Fatal("timeout")
	}
}
