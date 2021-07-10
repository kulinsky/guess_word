package main

import (
	"context"
	"log"

	"go.uber.org/zap"

	repository "github.com/kulinsky/guess_word/implementation/memory.repository"

	"github.com/kulinsky/guess_word/infrastructure"
	"github.com/kulinsky/guess_word/usecase"
)

func main() {
	logger, _ := zap.NewProduction()

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(logger)

	ucHandler := usecase.HandlerConstructor{
		Repository: repository.CreateRepository(),
		Logger:     logger,
	}.New()

	words := []string{"this", "is", "modern", "scalable", "high", "performance", "application"}
	attemptCount := 5

	ctx := context.Background()

	app := infrastructure.NewConsoleApp(attemptCount, words, ucHandler, logger)
	app.Start(ctx)
}
