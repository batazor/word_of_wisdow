package main

import (
	"context"
	"fmt"

	"github.com/batazor/word_of_wisdom/internal/pkg/logger"
	"github.com/batazor/word_of_wisdom/internal/pkg/tcp"
	repository "github.com/batazor/word_of_wisdom/internal/repository/bookofwisdom"
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

	// Init repository
	quotesRepository, err := repository.New("./internal/repository/bookofwisdom/data.json")
	if err != nil {
		panic(err)
	}

	// read data from the server
	for msg := range server.ReadCh {
		logger.Info("msg", zap.String("msg", string(msg)))

		// get random quote
		item, err := quotesRepository.GetRandomItem()
		if err != nil {
			panic(err)
		}

		replyMessage := fmt.Sprintf("%s\n", item.Quote)
		server.Send([]byte(replyMessage))
	}

	// TODO: Graceful shutdown
	return
}
