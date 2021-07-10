package domain

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var gameID = GenerateGameID()
var wordID = GenerateWordID()
var anotherWordID = GenerateWordID()
var attemptCount = 5
var word = NewWord(wordID, "hello")
var anotherWord = NewWord(anotherWordID, "world")

func TestNewGame(t *testing.T) {
	t.Run("normally create game", func(t *testing.T) {

		game, err := NewGame(gameID, &word, attemptCount)

		assert.Equal(t, err, nil)
		assert.Equal(t, game.ID, gameID)
	})

	t.Run("nil word", func(t *testing.T) {

		_, err := NewGame(gameID, nil, attemptCount)

		assert.Equal(t, err, ErrEmptyWord)
	})

	t.Run("negative attempt count", func(t *testing.T) {

		_, err := NewGame(gameID, &word, -1)

		assert.Equal(t, err, ErrNegativeAttemptCount)
	})
}

func TestUpdateGame(t *testing.T) {
	t.Run("normally update game", func(t *testing.T) {

		game, err := NewGame(gameID, &word, attemptCount)
		assert.Equal(t, err, nil)

		err = UpdateGame(game, SetAttemptCount(10))
		assert.Equal(t, err, nil)
		assert.Equal(t, game.AttemptCount, 10)
	})

	t.Run("normally update word", func(t *testing.T) {
		game, err := NewGame(gameID, &word, attemptCount)
		assert.Equal(t, err, nil)

		err = UpdateGame(game, SetWord(&anotherWord))
		assert.Equal(t, err, nil)
		assert.Equal(t, game.Word.String(), anotherWord.String())
	})
}
