package domain

type Word struct {
	ID    WordID
	value string
}

func NewWord(id WordID, s string) Word {
	return Word{
		ID:    id,
		value: s,
	}
}

func (w *Word) String() string {
	return w.value
}
