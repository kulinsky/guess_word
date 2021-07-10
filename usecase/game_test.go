package usecase

import (
	"context"
	"github.com/go-playground/assert/v2"
	"github.com/kulinsky/guess_word/domain"
	"github.com/kulinsky/guess_word/mocks"
	"github.com/kulinsky/guess_word/testdata"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestInteractor_GetRandomWord(t *testing.T) {

	repo := mocks.GameRepository{}

	repo.On("GetWordList", mock.AnythingOfType("*context.emptyCtx")).Return(
		func(ctx context.Context) []*domain.Word {
			return []*domain.Word{&testdata.FirstWord}
		},
		func(ctx context.Context) error {
			return nil
		},
	)

	ucHandler := HandlerConstructor{
		Repository: &repo,
		Logger:     zaptest.NewLogger(t),
	}.New()

	t.Run("get normally", func(t *testing.T) {
		w, err := ucHandler.GetRandomWord(context.Background())

		assert.Equal(t, err, nil)
		assert.Equal(t, w.ID, testdata.FWID)
	})
}
