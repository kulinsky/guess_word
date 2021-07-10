package domain

import (
	"strings"
	"time"
)

type GameStat struct {
	GameID       GameID
	AttemptCount int
	CurrentWord  string
	Win          bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Game struct {
	ID           GameID
	AttemptCount int
	WordID       *WordID
	Guessed      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Word         *Word // domain field
}

// NewGame фабричный метод, создает новую игру
func NewGame(gameID GameID, w *Word, attemptCount int) (*Game, error) {
	game := &Game{
		ID:        gameID,
		WordID:    &w.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// сначала нужно установить слово
	if err := UpdateGame(game, SetWord(w)); err != nil {
		return nil, err
	}

	guessed := make([]string, len(w.String()))

	for idx, _ := range guessed {
		guessed[idx] = "*"
	}

	if err := UpdateGame(
		game,
		SetAttemptCount(attemptCount),
		SetGuessed(strings.Join(guessed[:], "")),
		SetWord(w),
	); err != nil {
		return nil, err
	}

	return game, nil
}

// UpdateGame обновляет Game, принимает на вход функции обновления
func UpdateGame(initial *Game, opts ...func(fields *Game) error) error {
	for _, v := range opts {
		if err := v(initial); err != nil {
			return err
		}
	}

	initial.UpdatedAt = time.Now().UTC()

	return nil
}

// SetAttemptCount функция обновления для количество попыток, если на вход < 0 то возвращает ошибку
func SetAttemptCount(input int) func(fields *Game) error {
	return func(initial *Game) error {
		if input < 0 {
			return ErrNegativeAttemptCount
		}

		initial.AttemptCount = input

		return nil
	}
}

// SetGuessed функция обновления для слова
func SetGuessed(input string) func(fields *Game) error {
	return func(initial *Game) error {
		if input == "" {
			return ErrEmptyGuessed
		}

		if len(input) != len(initial.Word.String()) {
			return ErrDifferentLength
		}

		initial.Guessed = input

		return nil
	}
}

func SetWord(input *Word) func(fields *Game) error {
	return func(initial *Game) error {
		if input == nil {
			return ErrEmptyWord
		}

		initial.Word = input
		initial.WordID = &input.ID

		return nil
	}
}

// ReduceAttemptCount метод для уменьшения количеста попыток на 1
func (g *Game) ReduceAttemptCount() error {
	if err := UpdateGame(g, SetAttemptCount(g.AttemptCount-1)); err != nil {
		return err
	}

	return nil
}

// CheckLetter метод для проверки наличая буквы в искомом слове
func (g *Game) CheckLetter(s string) (int, error) {
	idx := -1

	if len(s) != 1 {
		return idx, ErrNotLetter
	}

	idx = strings.Index(g.Word.String(), s)

	return idx, nil
}

// SetGuessedLetter метод для обновления текущего состояния слова
func (g *Game) SetGuessedLetter(s string) (bool, error) {
	updated := false

	if len(s) != 1 {
		return false, ErrNotLetter
	}

	word := strings.Split(g.Guessed, "")

	for idx, v := range strings.Split(g.Word.String(), "") {
		if v == s {
			word[idx] = s

			updated = true
		}
	}

	if updated {
		if err := UpdateGame(g, SetGuessed(strings.Join(word[:], ""))); err != nil {
			return false, err
		}
	}

	return updated, nil
}

func (g *Game) IsWin() bool {
	switch g.Guessed {
	case g.Word.String():
		return true
	default:
		return false
	}
}

func (g *Game) GetGameStat() GameStat {
	return GameStat{
		GameID:       g.ID,
		AttemptCount: g.AttemptCount,
		CurrentWord:  g.Guessed,
		Win:          g.IsWin(),
		CreatedAt:    g.CreatedAt,
		UpdatedAt:    g.UpdatedAt,
	}
}
