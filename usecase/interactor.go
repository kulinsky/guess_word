package usecase

import (
	"context"
	"github.com/kulinsky/guess_word/domain"
	"go.uber.org/zap"
)

type interactor struct {
	repository GameRepository
	logger     *zap.Logger
}

type GameRepository interface {
	GameRW
	WordRW
}

type GameRW interface {
	SaveGame(ctx context.Context, game *domain.Game) error
	GetGame(ctx context.Context, gameID *domain.GameID) (*domain.Game, error)
}

type WordRW interface {
	SaveWord(ctx context.Context, word *domain.Word) error
	GetWordList(ctx context.Context) ([]*domain.Word, error)
}
