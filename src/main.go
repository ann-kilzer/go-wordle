package main

import (
	"fmt"

	"github.com/ann-kilzer/go-wordle/game"
)

func main() {
	fmt.Println("WORDLE")
	fmt.Println()
	fmt.Println("<x>: unused letter")
	fmt.Println("[X]: found letter")
	fmt.Println(" _ : used letter")
	fmt.Println()
	fmt.Println("[X]: The letter X is in the word and in the correct spot")
	fmt.Println("?x?: The letter X is in the word but in the wrong spot")
	fmt.Println("|_|: the letter is not in the word in any spot")
	fmt.Println()

	game := game.NewGame("BIRDS")
	game.Play()
}
