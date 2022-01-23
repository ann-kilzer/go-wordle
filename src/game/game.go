package game

import (
	"bufio"
	"fmt"
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

func (g *Game) isWin() bool {
	return g.currentRound().guess == g.word
}

func (g *Game) Play() error {
	for i := 0; i < ROUNDS; i++ {
		g.printRow()
		g.printLetters()

		err := g.readGuess()
		if err != nil {
			return err
		}

		fmt.Println(g.currentRound().guess)
		if g.isWin() {
			fmt.Println("Win!!!")
			break
			// todo
		}

		g.round += 1

	}

	return nil
}
