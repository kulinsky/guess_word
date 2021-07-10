package memrepo

import (
	"github.com/kulinsky/guess_word/usecase"
	"sync"
)

type rw struct {
	store *sync.Map
}

func CreateRepository() usecase.GameRepository {
	return rw{
		store: &sync.Map{},
	}
}
