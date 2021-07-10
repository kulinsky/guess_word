package usecase

import (
	"context"
	"log"

	"go.uber.org/zap"

	"github.com/kulinsky/guess_word/domain"
)

type GameLogic interface {
	StartNewGame(ctx context.Context, attemptCount int, w domain.Word) (*domain.GameID, error)
	GetGameStat(ctx context.Context, gameID *domain.GameID) (*domain.GameStat, error)
	Guess(ctx context.Context, gameID *domain.GameID, letter string) (bool, error)
}

type WordLogic interface {
	WordCreate(ctx context.Context, s string) (*domain.Word, error)
	GetRandomWord(ctx context.Context) (*domain.Word, error)
}

type Handler interface {
	GameLogic
	WordLogic
}

type HandlerConstructor struct {
	Repository GameRepository
	Logger     *zap.Logger
}

func (c HandlerConstructor) New() Handler {
	if c.Repository == nil {
		log.Fatal("Repository not set!")
	}

	return interactor{
		repository: c.Repository,
		logger:     c.Logger,
	}
}
