package main

import (
	"context"
	"go.uber.org/zap"

	repository "github.com/kulinsky/guess_word/implementation/memory.repository"

	"github.com/kulinsky/guess_word/infrastructure"
	"github.com/kulinsky/guess_word/usecase"
)

func main() {
	logger, _ := zap.NewProduction()

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	ucHandler := usecase.HandlerConstructor{
		Repository: repository.CreateRepository(),
		Logger:     logger,
	}.New()

	attemptCount := 5

	ctx := context.Background()

	app := infrastructure.NewConsoleApp(attemptCount, ucHandler, logger)

	// we need to init words, coz we use inmemory repo
	words := []string{"this", "is", "modern", "scalable", "high", "performance", "application"}
	app.InitWords(ctx, words)

	app.Start(ctx)
}
