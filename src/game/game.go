package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ann-kilzer/go-wordle/common"
	"github.com/ann-kilzer/go-wordle/dictionary"
	"github.com/ann-kilzer/go-wordle/game/answer"
	"github.com/ann-kilzer/go-wordle/game/round"
	"github.com/ann-kilzer/go-wordle/keyboard"
)

type Game struct {
	word         answer.Word // the answer
	reader       *bufio.Reader
	keyboard     *keyboard.Keyboard
	round        int
	rounds       [common.ROUNDS]*round.Round
	validGuesses *dictionary.ValidGuesses
}

func NewGame(word string, validGuesses *dictionary.ValidGuesses) *Game {
	return &Game{
		reader:       bufio.NewReader(os.Stdin),
		word:         answer.NewWord(word),
		validGuesses: validGuesses,
		keyboard:     keyboard.NewKeyboard(),
	}
}

func (g *Game) isWin() bool {
	return g.word.IsWin(g.currentRound().Guess)
}

// currentRound returns the current Round of the game, (or instantiates a new one when nil)
func (g *Game) currentRound() *round.Round {
	if g.rounds[g.round] == nil {
		g.rounds[g.round] = round.NewRound()
	}
	return g.rounds[g.round]
}

func (g *Game) Play() error {
	// print blanks for the 0th round
	g.printLetters()
	g.printResponse()

	for i := 0; i < common.ROUNDS; i++ {
		err := g.readGuess(i+1, common.ROUNDS)
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
