package infrastructure

import (
	"context"
	"github.com/kulinsky/guess_word/implementation/console"
	"github.com/kulinsky/guess_word/usecase"
	"go.uber.org/zap"
)

type ConsoleApp struct {
	AttemptCount   int
	InitialWords   []string
	CommandHandler console.CommandHandler
}

func NewConsoleApp(attemptCount int, words []string, uc usecase.Handler, logger *zap.Logger) ConsoleApp {
	ch := console.NewCommandHandler(uc, logger)

	return ConsoleApp{
		AttemptCount:   attemptCount,
		InitialWords:   words,
		CommandHandler: ch,
	}
}

func (app *ConsoleApp) Start(ctx context.Context) {
	app.CommandHandler.InitWords(ctx, app.InitialWords)
	app.CommandHandler.Start(ctx, app.AttemptCount)
}
