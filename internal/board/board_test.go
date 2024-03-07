package board_test

import (
	"github.com/mbobrovskyi/game-of-life/internal/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_New(t *testing.T) {
	b := board.NewWithGliderPattern(2, 2)
	expected := make([][]bool, 2)
	for i := 0; i < len(expected); i++ {
		expected[i] = make([]bool, 2)
	}
	assert.Equal(t, board.Board(expected), b)
}

func TestBoard_NewWithGlitterPatternEmpty(t *testing.T) {
	b := board.NewWithGliderPattern(2, 2)
	expected := board.New(2, 2)
	assert.Equal(t, expected, b)
}

func TestBoard_NewWithGlitterPattern(t *testing.T) {
	b := board.NewWithGliderPattern(5, 5)

	// _ _ _ _ _
	// _ _ X _ _
	// _ _ _ X _
	// _ X X X _
	// _ _ _ _ _
	expected := board.New(5, 5)
	expected[1][2] = true
	expected[2][3] = true
	expected[3][1] = true
	expected[3][2] = true
	expected[3][3] = true

	assert.Equal(t, expected, b)
}

func TestBoard_Set(t *testing.T) {
	b := board.New(0, 0)
	assert.Error(t, b.Set(1, 1, true))
	assert.Equal(t, board.OutOfRangeError, b.Set(1, 1, true))
}

func TestBoard_Neighbors(t *testing.T) {
	b := board.NewWithGliderPattern(3, 3)

	tests := []struct {
		y        int
		x        int
		expected int
	}{
		{0, 0, 1},
		{0, 1, 1},
		{0, 2, 2},
		{1, 0, 3},
		{1, 1, 5},
		{1, 2, 3},
		{2, 0, 1},
		{2, 1, 3},
		{2, 2, 2},
	}

	for _, test := range tests {
		neighbors, err := b.Neighbors(test.y, test.x)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, neighbors)
	}
}

func TestBoard_NeighborsWithError(t *testing.T) {
	b := board.New(0, 0)
	_, err := b.Neighbors(1, 1)
	assert.Error(t, err)
	assert.Equal(t, board.OutOfRangeError, err)
}

func TestGame_Next(t *testing.T) {
	b := board.NewWithGliderPattern(5, 5)
	b.Print()

	next := b.Next()
	next.Print()

	// _ _ _ _ _
	// _ _ _ _ _
	// _ X _ X _
	// _ _ X X _
	// _ _ X _ _
	expected := board.New(5, 5)
	expected[2][1] = true
	expected[2][3] = true
	expected[3][2] = true
	expected[3][3] = true
	expected[4][2] = true

	assert.Equal(t, expected, next)
}
