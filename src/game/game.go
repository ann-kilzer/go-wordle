package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ann-kilzer/go-wordle/dictionary"
	"github.com/ann-kilzer/go-wordle/keyboard"
)

const WORD_LENGTH = 5
const ROUNDS = 6

type Game struct {
	word         Word // the answer
	reader       *bufio.Reader
	keyboard     *keyboard.Keyboard
	round        int
	rounds       [ROUNDS]*Round
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
	return g.word.isWin(g.currentRound().guess)
}

// markUsedLetters updates which letters have been used on the letterRecord "keyboard"
func (g *Game) markUsedLetters() {
	guess := g.currentRound().guess
	eval := g.currentRound().eval
	for i := 0; i < len(eval); i++ {
		if eval[i] == GREEN {
			g.keyboard.MarkMatch(guess[i])
		} else if eval[i] == YELLOW {
			g.keyboard.MarkMatch(guess[i])
		} else {
			g.keyboard.MarkNoMatch(guess[i])
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
