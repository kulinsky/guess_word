package infrastructure

import (
	"context"
	"github.com/kulinsky/guess_word/implementation/console"
	"github.com/kulinsky/guess_word/usecase"
	"go.uber.org/zap"
)

type ConsoleApp struct {
	AttemptCount   int
	CommandHandler console.CommandHandler
}

func NewConsoleApp(attemptCount int, uc usecase.Handler, logger *zap.Logger) ConsoleApp {
	ch := console.NewCommandHandler(uc, logger)

	return ConsoleApp{
		AttemptCount:   attemptCount,
		CommandHandler: ch,
	}
}

func (app *ConsoleApp) Start(ctx context.Context) {
	app.CommandHandler.Start(ctx, app.AttemptCount)
}

func (app *ConsoleApp) InitWords(ctx context.Context, words []string) {
	app.CommandHandler.InitWords(ctx, words)
}
