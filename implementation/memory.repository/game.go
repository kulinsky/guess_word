package memrepo

import (
	"context"
	"fmt"

	"github.com/kulinsky/guess_word/domain"
)

func (rw rw) SaveGame(ctx context.Context, game *domain.Game) error {
	gameID := fmt.Sprint("Games:", game.ID)
	rw.store.Store(gameID, *game)

	return nil
}

func (rw rw) GetGame(ctx context.Context, gameID *domain.GameID) (*domain.Game, error) {
	recID := fmt.Sprint("Games:", gameID)
	value, ok := rw.store.Load(recID)

	if !ok {
		return nil, domain.ErrNotFound
	}

	game, ok := value.(domain.Game)
	if !ok {
		return nil, domain.ErrNotGame
	}

	return &game, nil
}
