package game

import (
	"bufio"
	"os"
)

var word = "BIRDS"

const WORD_LENGTH = 5
const ROUNDS = 6

type Game struct {
	reader      *bufio.Reader
	usedLetters [26]bool
	round       int
	rounds      [ROUNDS]*Round
}

func NewGame() *Game {
	return &Game{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (g *Game) Play() {
	for i := 0; i < ROUNDS; i++ {
		g.printRow()
		g.printLetters()
		g.readGuess()
		g.round += 1
	}
}
