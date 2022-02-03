package game

import (
	"fmt"
	"strings"

	"github.com/ann-kilzer/go-wordle/round"
)

func (g *Game) currentRound() *round.Round {
	if g.rounds[g.round] == nil {
		g.rounds[g.round] = round.NewRound()
	}
	return g.rounds[g.round]
}

func (g *Game) printLetters() {
	fmt.Println(g.keyboard)
}

// |_| not found
// ?x? wrong location
// [X] found
func (g *Game) printResponse() {
	guess := g.currentRound().Guess
	eval := g.currentRound().Eval
	for i := 0; i < len(eval); i++ {
		letter := string(guess[i])
		switch eval[i] {
		case GREEN:
			fmt.Printf("[%v]", letter)
		case YELLOW:
			fmt.Printf("?%v?", strings.ToLower(letter))
		case BLACK:
			fmt.Print("|_|")
		}
	}

	fmt.Println()
}

func (g *Game) readGuess(r, rounds int) (err error) {
	var guess string

	for {
		fmt.Printf("%d/%d\n>", r, rounds)

		guess, err = g.reader.ReadString('\n')
		if err != nil {
			return err
		}
		guess = strings.ToUpper(strings.TrimSpace(guess))

		if len(guess) != round.WORD_LENGTH {
			fmt.Printf("Invalid length. Please enter %v letters\n", round.WORD_LENGTH)
		} else if g.validGuesses.Contains(guess) {
			break
		} else {
			fmt.Println("Invalid word")
		}
	}

	g.currentRound().SetGuess(guess)

	return nil
}

// evaluateRound must be called after the guess is read
func (g *Game) evaluateRound() error {
	// evaluate the guess
	eval, err := g.word.evaluateGuess(g.currentRound().Guess)
	if err != nil {
		return err
	}

	g.currentRound().SetEval(eval)
	return nil
}
