package memrepo

import (
	"context"
	"fmt"
	"github.com/kulinsky/guess_word/domain"
	"strings"
)

func (rw rw) SaveWord(ctx context.Context, word *domain.Word) error {
	courseID := fmt.Sprint("Words:", word.ID)
	rw.store.Store(courseID, *word)

	return nil
}

func (rw rw) GetWordList(ctx context.Context) ([]*domain.Word, error) {
	var res []*domain.Word

	rw.store.Range(func(key, value interface{}) bool {
		if !strings.HasPrefix(fmt.Sprintf("%v", key), "Words:") {
			return true
		}
		w, ok := value.(domain.Word)
		if !ok {
			return true
		}
		res = append(res, &w)

		return true
	})

	return res, nil
}
