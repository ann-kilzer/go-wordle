package game

import (
	"fmt"
	"strings"
)

const UPPER_A = 65
const LOWER_A = 97
const EMPTY_GUESS = "     " // 5 blank spaces

type Round struct {
	guess    string // user input
	feedback string // the response to the user
}

// setGuess records the guess in the round as an uppercase string
func (r *Round) setGuess(guess string) {
	r.guess = strings.ToUpper(guess)
}

func (g *Game) currentRound() *Round {
	if g.rounds[g.round] == nil {
		g.rounds[g.round] = &Round{guess: EMPTY_GUESS}
	}
	return g.rounds[g.round]
}

// TODO: QWERTY order
func (g *Game) printLetters() {
	// TODO
	for i := 0; i < len(g.letterRecord); i++ {
		switch g.letterRecord[i] {
		case UNUSED:
			fmt.Printf("<%c>", i+LOWER_A)
		case MATCH:
			fmt.Printf("[%c]", i+UPPER_A)
		case NO_MATCH:
			fmt.Print(" _ ")
		}
	}

	fmt.Println()
}

// |_| not found
// ?x? wrong location
// [X] found
func (g *Game) printResponse() {
	guess := g.currentRound().guess
	output := g.word.evaluateGuess(guess)
	for i := 0; i < len(output); i++ {
		letter := string(guess[i])
		switch output[i] {
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
		guess = strings.TrimSpace(guess)

		if len(guess) == WORD_LENGTH {
			break
		} else {
			fmt.Printf("Invalid length. Please enter %v letters\n", WORD_LENGTH)
		}
	}

	g.currentRound().setGuess(guess)

	return nil
}
