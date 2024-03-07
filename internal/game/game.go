package game

import (
	"github.com/mbobrovskyi/game-of-life/internal/board"
	"time"
)

type Game struct {
	board board.Board
}

func New(height, width int) *Game {
	g := &Game{
		board: board.NewWithGliderPattern(height, width),
	}

	return g
}

func (g *Game) Start() {
	for {
		g.board.Print()
		time.Sleep(time.Second)
		g.board = g.board.Next()
	}
}
