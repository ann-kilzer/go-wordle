package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const WORD_LENGTH = 5
const ROUNDS = 6

// for letterRecord
const UNUSED = 0
const NO_MATCH = 1
const MATCH = 2

type Game struct {
	word         string
	reader       *bufio.Reader
	letterRecord [26]int // track used and found letters
	round        int
	rounds       [ROUNDS]*Round
}

func NewGame(word string) *Game {
	return &Game{
		reader: bufio.NewReader(os.Stdin),
		word:   word,
	}
}

func (g *Game) inWord(letter string) bool {
	return strings.Contains(g.word, letter)
}

func (g *Game) isWin() bool {
	return g.currentRound().guess == g.word
}

func (g *Game) markUsedLetters() {
	for i := 0; i < WORD_LENGTH; i++ {
		letter := string(g.currentRound().guess[i])
		index := g.currentRound().guess[i] - UPPER_A
		if g.inWord(letter) {
			g.letterRecord[index] = MATCH
		} else {
			g.letterRecord[index] = NO_MATCH
		}
	}
}

func (g *Game) Play() error {
	// print blanks for the 0th round
	g.printLetters()
	g.printRow()

	for i := 0; i < ROUNDS; i++ {
		err := g.readGuess()
		if err != nil {
			return err
		}

		g.markUsedLetters()

		g.printLetters()
		g.printRow()
		if g.isWin() {
			fmt.Println("Win!!!")
			break
			// todo
		}

		g.round += 1

	}

	return nil
}
