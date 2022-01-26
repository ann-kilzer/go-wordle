package game

import (
	"bufio"
	"fmt"
	"os"
)

const WORD_LENGTH = 5
const ROUNDS = 6

// for letterRecord
const UNUSED = 0
const NO_MATCH = 1
const MATCH = 2

type Game struct {
	word         Word // the answer
	reader       *bufio.Reader
	letterRecord [26]int // track used and found letters
	round        int
	rounds       [ROUNDS]*Round
}

func NewGame(word string) *Game {
	return &Game{
		reader: bufio.NewReader(os.Stdin),
		word:   NewWord(word),
	}
}

func (g *Game) isWin() bool {
	return g.word.isWin(g.currentRound().guess)
}

// todo: refactor this
func (g *Game) markUsedLetters() {
	for i := 0; i < WORD_LENGTH; i++ {
		guess := g.currentRound().guess
		index := guess[i] - UPPER_A // ascii index
		if g.word.isGreen(guess, i) {
			g.letterRecord[index] = MATCH
		} else if g.word.isYellow(guess, i) {
			g.letterRecord[index] = MATCH
		} else {
			g.letterRecord[index] = NO_MATCH
		}
	}
}

func (g *Game) Play() error {
	// print blanks for the 0th round
	g.printLetters()
	g.printResponse()

	for i := 0; i < ROUNDS; i++ {
		err := g.readGuess()
		if err != nil {
			return err
		}

		g.markUsedLetters()
		g.printLetters()

		g.printResponse()
		if g.isWin() {
			fmt.Println("Win!!!")
			break
			// todo
		}

		g.round += 1

	}

	return nil
}
