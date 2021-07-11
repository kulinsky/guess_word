package formatter

import (
	"github.com/kulinsky/guess_word/domain"
	"time"
)

type GameResponse struct {
	ID           domain.GameID  `json:"id"`
	AttemptCount int            `json:"attempt_count"`
	WordID       *domain.WordID `json:"word_id"`
	Guessed      string         `json:"guessed"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func NewGameResponseFromDomain(game *domain.Game) *GameResponse {
	return &GameResponse{
		ID:           game.ID,
		AttemptCount: game.AttemptCount,
		WordID:       game.WordID,
		Guessed:      game.Guessed,
		CreatedAt:    game.CreatedAt,
		UpdatedAt:    game.UpdatedAt,
	}
}

type GameStatResponse struct {
	GameID       domain.GameID `json:"game_id"`
	AttemptCount int           `json:"attempt_count"`
	CurrentWord  string        `json:"current_word"`
	Win          bool          `json:"win"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func NewGameStatResponseFromDomain(stat *domain.GameStat) *GameStatResponse {
	return &GameStatResponse{
		GameID:       stat.GameID,
		AttemptCount: stat.AttemptCount,
		CurrentWord:  stat.CurrentWord,
		Win:          stat.Win,
		CreatedAt:    stat.CreatedAt,
		UpdatedAt:    stat.UpdatedAt,
	}
}
