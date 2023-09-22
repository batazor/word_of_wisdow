package main

import (
	"context"
	"fmt"

	"github.com/batazor/word_of_wisdom/internal/pkg/logger"
	"github.com/batazor/word_of_wisdom/internal/pkg/tcp"
)

func main() {
	// Init context
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init logger
	log, err := logger.New()
	if err != nil {
		panic(err)
	}

	// Init client
	uri := "localhost:8080"
	client, err := tcp.NewClient(uri)
	if err != nil {
		panic(err)
	}

	// Send ping
	_, err = client.Conn.Write([]byte("ping"))

	// Read message
	for msg := range client.ReadCh {
		log.Info(fmt.Sprintf("Message: %s", msg))
	}

	// TODO: Graceful shutdown
	return
}
