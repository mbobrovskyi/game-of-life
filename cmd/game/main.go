package main

import (
	"github.com/mbobrovskyi/game-of-life/internal/game"
)

const (
	height = 25
	width  = 25
)

func main() {
	g := game.New(height, width)
	g.Start()
}
