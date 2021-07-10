package usecase

import (
	"context"
	"math/rand"
	"time"

	"github.com/kulinsky/guess_word/domain"
)

func (i interactor) WordCreate(ctx context.Context, s string) (*domain.Word, error) {
	word := domain.NewWord(domain.GenerateWordID(), s)

	if err := i.repository.SaveWord(ctx, &word); err != nil {
		i.logger.Error(err.Error())

		return nil, domain.ErrTechnical
	}

	return &word, nil
}

func (i interactor) GetRandomWord(ctx context.Context) (*domain.Word, error) {
	words, err := i.repository.GetWordList(ctx)

	if err != nil {
		i.logger.Error(err.Error())

		return nil, domain.ErrTechnical
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	return words[r.Intn(len(words))], nil
}
