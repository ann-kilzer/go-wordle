package main

import (
	"fmt"
)

func main() {
	fmt.Println("WORDLE")
	fmt.Println()

	game := game.NewGame()
	game.Play()
}
