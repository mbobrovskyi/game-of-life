package board

import (
	"fmt"
	"github.com/mbobrovskyi/game-of-life/internal/terminal"
)

var OutOfRangeError = fmt.Errorf("out of range error")

type Board [][]bool

func New(height, width int) Board {
	b := make(Board, height)

	for y := 0; y < len(b); y++ {
		b[y] = make([]bool, width)
	}

	return b
}

// NewWithGliderPattern
// _ _ _ _ _
// _ _ X _ _
// _ _ _ X _
// _ X X X _
// _ _ _ _ _
func NewWithGliderPattern(height, width int) Board {
	b := New(height, width)

	if len(b) >= 3 || len(b[0]) >= 3 {
		heightMid := len(b) / 2
		widthMid := len(b[0]) / 2

		b[heightMid-1][widthMid] = true
		b[heightMid][widthMid+1] = true
		b[heightMid+1][widthMid-1] = true
		b[heightMid+1][widthMid] = true
		b[heightMid+1][widthMid+1] = true
	}

	return b
}

func (b Board) isOutOfRange(y, x int) bool {
	return len(b) < y || len(b[y]) < x
}

func (b Board) Set(y, x int, value bool) error {
	if b.isOutOfRange(y, x) {
		return OutOfRangeError
	}

	b[y][x] = value

	return nil
}

func (b Board) Neighbors(y, x int) (int, error) {
	if b.isOutOfRange(y, x) {
		return 0, OutOfRangeError
	}

	var count int

	if y > 0 {
		if x > 0 && b[y-1][x-1] {
			count++
		}

		if b[y-1][x] {
			count++
		}

		if x < len(b[y-1])-1 && b[y-1][x+1] {
			count++
		}
	}

	if x > 0 && b[y][x-1] {
		count++
	}

	if x < len(b[y])-1 && b[y][x+1] {
		count++
	}

	if y < len(b)-1 {
		if x > 0 && b[y+1][x-1] {
			count++
		}

		if b[y+1][x] {
			count++
		}

		if x < len(b[y+1])-1 && b[y+1][x+1] {
			count++
		}
	}

	return count, nil
}

func (b Board) Next() Board {
	if len(b) == 0 {
		return b
	}

	next := New(len(b), len(b[0]))

	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[y]); x++ {
			neighbors, _ := b.Neighbors(y, x)

			if neighbors < 2 || neighbors > 3 {
				next[y][x] = false
			} else if neighbors == 2 {
				next[y][x] = b[y][x]
			} else if neighbors == 3 {
				next[y][x] = true
			}

		}
	}

	return next
}

func (b Board) Print() {
	terminal.Clear()

	fmt.Println("Game of Life")

	for y := 0; y < len(b); y++ {
		fmt.Print("|")

		for x := 0; x < len(b[y]); x++ {
			if b[y][x] {
				fmt.Print(" X ")
			} else {
				fmt.Print("   ")
			}
			fmt.Print("|")
		}

		fmt.Println()
	}
}
