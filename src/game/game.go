package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ann-kilzer/go-wordle/dictionary"
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
	validGuesses *dictionary.ValidGuesses
}

func NewGame(word string, validGuesses *dictionary.ValidGuesses) *Game {
	return &Game{
		reader:       bufio.NewReader(os.Stdin),
		word:         NewWord(word),
		validGuesses: validGuesses,
	}
}

func (g *Game) isWin() bool {
	return g.word.isWin(g.currentRound().guess)
}

// markUsedLetters updates which letters have been used on the letterRecord "keyboard"
func (g *Game) markUsedLetters() {
	guess := g.currentRound().guess
	eval := g.currentRound().eval
	for i := 0; i < len(eval); i++ {
		index := guess[i] - UPPER_A        // ascii index
		if g.letterRecord[index] > BLACK { // already yellow or green
			continue
		}
		if eval[i] == GREEN {
			g.letterRecord[index] = MATCH
		} else if eval[i] == YELLOW {
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
			return nil
		}

		g.round += 1
	}

	fmt.Println("Sorry! The word was ", g.word)

	return nil
}
