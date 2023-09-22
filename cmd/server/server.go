package main

import (
	"context"

	"github.com/batazor/word_of_wisdom/internal/pkg/logger"
	"github.com/batazor/word_of_wisdom/internal/pkg/tcp"
	"go.uber.org/zap"
)

func main() {
	// Init context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init logger
	logger, err := logger.New()
	if err != nil {
		panic(err)
	}

	// Init server
	uri := "localhost:8080"
	server, err := tcp.NewServer(ctx, uri, logger)
	if err != nil {
		panic(err)
	}

	// read data from the server
	for msg := range server.ReadCh {
		logger.Info("msg", zap.String("msg", string(msg)))
		server.Send([]byte("pong\n"))
	}

	// Graceful shutdown
	return
}
