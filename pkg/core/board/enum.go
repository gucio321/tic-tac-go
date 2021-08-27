package board

import "log"

// BoardW, BoardH are board's width and height.
const (
	BaseBoardW   = 3
	BaseBoardH   = 3
	BaseChainLen = 3
)

// GetWinBoard returns winning indexes list.
func (b *Board) GetWinBoard(l int) [][]int {
	w, h := b.Width(), b.Height()
	// for w = h:
	// n = (w-l+1)*h + (h-l+1) * w + 2 * ((w or h)-l+1)
	// generally (if s = w = h) n = (s-l+1)*s + (s-l+1) *w + 2 * (s - l + 1)
	winningIndexes := make([][]int, 0)

	// horizontal indexes
	for row := 0; row < h; row++ {
		for rowIdx := 0; rowIdx+l <= w; rowIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = row*w + rowIdx + idx
			}

			winningIndexes = append(winningIndexes, line)
		}
	}

	// vertical indexes
	for col := 0; col < w; col++ {
		for colIdx := 0; colIdx+l <= h; colIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = col + colIdx*w + idx*w
			}

			winningIndexes = append(winningIndexes, line)
		}
	}

	for x := 0; x < h; x++ {
		for xIdx := 0; (x*w+xIdx*w+xIdx)+((l-1)*w+(l-1)) <= h*w-1; xIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = (x*w + xIdx) + (idx*w + idx)
			}

			winningIndexes = append(winningIndexes, line)
		}
	}

	for bx := 0; bx < h; bx++ {
		for bxIdx := 0; bx*w+(bxIdx+l)*w <= h*w; bxIdx++ {
			line := make([]int, l)
			for idx := range line {
				line[idx] = (bx * w) + (bxIdx+idx)*w + (l - idx - 1)
			}

			winningIndexes = append(winningIndexes, line)
		}
	}

	return winningIndexes
}

// GetCorners returns board's corners.
func (b *Board) GetCorners() (result []int) {
	w, h := b.Width(), b.Height()
	result = []int{
		0,           // upper left
		w - 1,       // upper right
		w * (h - 1), // botton left
		w*h - 1,     // botton right
	}

	return
}

// GetOppositeCorner returns a corner in an opposite to given.
func (b *Board) GetOppositeCorner(c int) int {
	corners := b.GetCorners()
	for n, corner := range corners {
		if corner == c {
			return corners[len(corners)-1-n]
		}
	}

	panic("Tic-Tac-Go: board.(*Board).GetOppositeCorner: invalid corner index. Did you given a corner index?")
}

// ConvertIndex converts index from smaller to larger board (fiction-width, fiction-height, real-width, real-height)
/*
+---+---+---+---+---+
| x | x | x | x | x |
+---+---+---+---+---+    +---+---+---+
| x | o | o | o | x |    | o | o | o |
+---+---+---+---+---+    +---+---+---+
| x | o | o | o | x | => | o | o | o |
+---+---+---+---+---+    +---+---+---+
| x | o | o | o | x |    | o | o | o |
+---+---+---+---+---+    +---+---+---+
| x | x | x | x | x |
+---+---+---+---+---+

                 +---+---+---+---+---+
                 | 0 | 1 | 2 | 3 | 4 |
+---+---+---+    +---+---+---+---+---+
| 0 | 1 | 2 |    | 5 | 6 | 7 | 8 | 9 |
+---+---+---+    +---+---+---+---+---+
| 3 | 4 | 5 | => |10 |11 |12 |13 |14 |
+---+---+---+    +---+---+---+---+---+
| 6 | 7 | 8 |    |15 |16 |17 |18 |19 |
+---+---+---+    +---+---+---+---+---+
		 |20 |21 |22 |23 |24 |
		 +---+---+---+---+---+
*/
func (b *Board) ConvertIndex(fw, fh, idx int) int {
	rw, rh := b.Width(), b.Height()
	// static checks: check if fiction size isn't greater than real
	if !(fh <= rh) || !(fw <= rw) {
		log.Fatal("invalid input: input should be: fh > rh || fw > rw")
	}

	// static check: check if one dimension isn't odd and second even number
	if fh%2 != rh%2 || fw%2 != rw%2 {
		log.Fatal("invalid input: cannot process even and odd numbers together")
	}

	// idx is a list index, we need to make it a... index in real bord starting from 1
	idx++

	// create result:
	/* example:
	+---+---+---+---+---+ rows containing "x" should be added
	| x | x | x | x | x | independently from idx.
	+---+---+---+---+---+
	| x | o | o | o | x | the number of indexes in this rows
	+---+---+---+---+---+ ("x" rows on top) is equal to the half
	| x | o | u | o | x | of a diffrence the big board's ("x") height
	+---+---+---+---+---+ and a small board's ("o") height multiplied by "o" board's width
	| x | o | o | o | x |
	+---+---+---+---+---+ here: 4 * [ ( 5-3) / 2 ] = 4 * (2/2) = 4
	| x | x | x | x | x |
	+---+---+---+---+---+ the result of this operation is the last index of "x" rows on top
	*/
	result := rw * ((rh - fh) / 2) // nolint:gomnd // half of real and fiction size diffrence

	for idx > 0 {
		// when idx - fiction board's width it means,
		// that we need to continue loop
		// example:
		// let's suppose, that input idx = 4
		// it means, that it is a 4 index on "o" board (see above - marked as "u")
		// in first loop iteration, 4 - 2 > 0, so
		// idx is decreased by 3 ("o"-board's width) and
		// our result is increased by real board's size (4)
		// so our result is equal to 8 for now
		//
		// if this condition is passed our loop starts again.
		if idx-fw > 0 {
			idx -= fw
			result += rw

			continue
		}

		/*
			if above condition isn't passed, it means, that
			the index, we're searching for is in current line
			in our example:
			we need to add a starting column index (one "x" index before "o"/"u" line in my draw above)
			like the first rows, they are equal to:
			(rw - fw) / 2 = (5 - 3)/2 = 2 / 2 = 1
			later, we just add current idx value and... tada!
		*/
		result += (rw - fw) / 2 // nolint:gomnd // half of real and fiction size diffrence

		result += idx

		break
	}

	// we must make it an list index (which starts from 0
	return result - 1
}

// GetSides returns sidde indexes of board's edges.
func (b *Board) GetSides() (result []int) {
	w, h := b.Width(), b.Height()
	for i := 1; i < w-1; i++ {
		result = append(result, i)
	}

	for i := 1; i < h-1; i++ {
		result = append(result, i*w, (i*w)+w-1)
	}

	for i := 1; i < w-1; i++ {
		result = append(result, h*(w-1)+i)
	}

	return result
}

// GetCenter returns bard center (if exists).
func (b *Board) GetCenter() []int {
	w, h := b.Width(), b.Height()
	if w%2 == 0 || h%2 == 0 {
		return []int{}
	}

	return []int{(h-1)/2*w + (w-1)/2}
}

// IsEdgeIndex returns true if i is an index on board edge.
func (b *Board) IsEdgeIndex(i int) bool {
	w, h := b.Width(), b.Height()
	if i-w < 0 {
		return true
	}

	for j := 1; j < h; j++ {
		if i == j*w || i == j*w+w-1 {
			return true
		}
	}

	return i >= (h-1)*w
}
