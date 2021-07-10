package domain

import "errors"

var (
	ErrNegativeAttemptCount = errors.New("attempt count can't be negative")
	ErrEmptyGuessed         = errors.New("guessed word can't be empty")
	ErrEmptyWord            = errors.New("word can't be empty")
	ErrDifferentLength      = errors.New("word and guessed has different length")
	ErrNotLetter            = errors.New("it is not a letter")
	ErrGameOver             = errors.New("gg, game over, you lost")
	ErrTechnical            = errors.New("a technical error happened")
	ErrNotFound             = errors.New("entity not found")
	ErrNotGame              = errors.New("not a game stored at key")
	ErrAlreadyBeen          = errors.New("the letter already been guessed")
)
