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

func TestTCP(t *testing.T) {
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

	// Set a timeout for receiving messages
	timeout := time.After(3 * time.Second)

	// Read data from the server
	select {
	case v := <-server.ReadCh:
		if string(v) != "ping" {
			t.Fatal("expected ping, got", string(v))
		}
	case <-timeout:
		t.Fatal("timeout")
	}
}
