package usecase

import (
	"context"
	"github.com/kulinsky/guess_word/domain"
	"strings"
)

func (i interactor) StartNewGame(ctx context.Context, attemptCount int, w *domain.Word) (*domain.GameID, error) {
	gameID := domain.GenerateGameID()

	game, err := domain.NewGame(gameID, w, attemptCount)

	if err != nil {
		i.logger.Error(err.Error())

		return nil, err
	}

	if err := i.repository.SaveGame(ctx, game); err != nil {
		i.logger.Error(err.Error())

		return nil, domain.ErrTechnical
	}

	return &gameID, nil
}

func (i interactor) GetGameStat(ctx context.Context, gameID *domain.GameID) (*domain.GameStat, error) {
	game, err := i.repository.GetGame(ctx, gameID)

	if err != nil {
		i.logger.Error(err.Error())

		return nil, domain.ErrTechnical
	}

	stat := game.GetGameStat()

	return &stat, nil
}

func (i interactor) Guess(ctx context.Context, gameID *domain.GameID, letter string) (bool, error) {
	game, err := i.repository.GetGame(ctx, gameID)

	if err != nil {
		i.logger.Error(err.Error())

		return false, domain.ErrTechnical
	}

	if game.AttemptCount < 1 && !game.IsWin() {
		return false, domain.ErrGameOver
	}

	if strings.Contains(game.Guessed, letter) {
		i.logger.Warn(domain.ErrAlreadyBeen.Error())

		return true, nil
	}

	guessed, err := game.CheckLetter(letter)

	if err != nil {
		i.logger.Error(err.Error())

		return false, err
	}

	if guessed {
		if _, err := game.SetGuessedLetter(letter); err != nil {
			return true, err
		}

		if err := i.repository.SaveGame(ctx, game); err != nil {
			i.logger.Error(err.Error())

			return true, domain.ErrTechnical
		}

		return true, nil
	} else {
		if err := game.ReduceAttemptCount(); err != nil {
			if err != domain.ErrNegativeAttemptCount {
				i.logger.Error(err.Error())

				return false, domain.ErrTechnical
			}

			return false, domain.ErrGameOver
		}

		if err := i.repository.SaveGame(ctx, game); err != nil {
			i.logger.Error(err.Error())

			return false, domain.ErrTechnical
		}

		return false, nil
	}
}
