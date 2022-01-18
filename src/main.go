package main

import (
	"fmt"

	"github.com/ann-kilzer/go-wordle/game"
)

func main() {
	fmt.Println("WORDLE")
	fmt.Println()

	game := game.NewGame("BIRDS")
	game.Play()
}
