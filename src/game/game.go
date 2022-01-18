package game

import (
	"bufio"
	"os"
)

const WORD_LENGTH = 5
const ROUNDS = 6

type Game struct {
	word        string
	reader      *bufio.Reader
	usedLetters [26]bool
	round       int
	rounds      [ROUNDS]*Round
}

func NewGame(word string) *Game {
	return &Game{
		reader: bufio.NewReader(os.Stdin),
		word:   word,
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
