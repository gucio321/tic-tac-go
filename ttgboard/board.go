package ttgboard

type board []*Letter

func newBoard(w, h int) *board {
	result := &board{}
	*result = make([]*Letter, w*h)

	for i := range *result {
		(*result)[i] = newBoardIndex()
	}

	return result
}

func (b *board) setIndexState(i int, state Letter) {
	(*b)[i].SetState(state)
}

func (b *board) getIndexState(i int) Letter {
	return *(*b)[i]
}

func (b *board) isIndexFree(i int) bool {
	return (*b)[i].IsNone()
}
