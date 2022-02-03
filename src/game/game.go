package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ann-kilzer/go-wordle/dictionary"
	"github.com/ann-kilzer/go-wordle/keyboard"
	"github.com/ann-kilzer/go-wordle/round"
)

const ROUNDS = 6

type Game struct {
	word         Word // the answer
	reader       *bufio.Reader
	keyboard     *keyboard.Keyboard
	round        int
	rounds       [ROUNDS]*round.Round
	validGuesses *dictionary.ValidGuesses
}

func NewGame(word string, validGuesses *dictionary.ValidGuesses) *Game {
	return &Game{
		reader:       bufio.NewReader(os.Stdin),
		word:         NewWord(word),
		validGuesses: validGuesses,
		keyboard:     keyboard.NewKeyboard(),
	}
}

func (g *Game) isWin() bool {
	return g.word.isWin(g.currentRound().Guess)
}

// markUsedLetters updates which letters have been used on the letterRecord "keyboard"
func (g *Game) markUsedLetters() {
	guess := g.currentRound().Guess
	eval := g.currentRound().Eval
	for i := 0; i < round.WORD_LENGTH; i++ {
		letterByte := guess[i]
		if eval[i] == GREEN {
			g.keyboard.MarkMatch(letterByte)
		} else if eval[i] == YELLOW {
			g.keyboard.MarkMatch(letterByte)
		} else {
			g.keyboard.MarkNoMatch(letterByte)
		}
	}
}

func (g *Game) Play() error {
	// print blanks for the 0th round
	g.printLetters()
	g.printResponse()

	for i := 0; i < ROUNDS; i++ {
		err := g.readGuess(i+1, ROUNDS)
		if err != nil {
			return err
		}

		err = g.evaluateRound()
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
