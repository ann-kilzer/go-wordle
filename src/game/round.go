package game

import (
	"fmt"
	"strings"
)

const UPPER_A = 65
const LOWER_A = 97
const EMPTY_GUESS = "     " // 5 blank spaces

type Round struct {
	guess string           // user input
	eval  [WORD_LENGTH]int // the evaluation of the guess
}

// setGuess records the guess in the round as an uppercase string
func (r *Round) setGuess(guess string) {
	r.guess = strings.ToUpper(guess)
}

// setEval records the evaluation in the round
func (r *Round) setEval(eval [WORD_LENGTH]int) {
	r.eval = eval
}

func (g *Game) currentRound() *Round {
	if g.rounds[g.round] == nil {
		g.rounds[g.round] = &Round{guess: EMPTY_GUESS}
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
	guess := g.currentRound().guess
	eval := g.word.evaluateGuess(guess)
	g.currentRound().setEval(eval)
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

func (g *Game) readGuess() (err error) {
	var guess string

	for {
		fmt.Print(">")

		guess, err = g.reader.ReadString('\n')
		if err != nil {
			return err
		}
		guess = strings.ToUpper(strings.TrimSpace(guess))

		if len(guess) != WORD_LENGTH {
			fmt.Printf("Invalid length. Please enter %v letters\n", WORD_LENGTH)
		} else if g.validGuesses.Contains(guess) {
			break
		} else {
			fmt.Println("Invalid word")
		}
	}

	g.currentRound().setGuess(guess)

	return nil
}
