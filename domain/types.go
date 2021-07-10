package domain

import "github.com/google/uuid"

type GameID = uuid.UUID

type WordID = uuid.UUID

func GenerateGameID() GameID {
	return uuid.New()
}

func GenerateWordID() WordID {
	return uuid.New()
}
