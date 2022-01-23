package game

import (
	"fmt"
	"strings"
)

const A = 65

type Round struct {
	guess    string // user input
	feedback string // the response to the user
}

func (r *Round) setGuess(guess string) {
	r.guess = guess
}

func (g *Game) currentRound() *Round {
	if g.rounds[g.round] == nil {
		g.rounds[g.round] = &Round{}
	}
	return g.rounds[g.round]
}

// TODO: QWERTY order
func (g *Game) printLetters() {
	// TODO
	for i := 0; i < len(g.usedLetters); i++ {
		if g.usedLetters[i] {
			fmt.Print("[ ]")
		} else {
			fmt.Printf("[%c]", i+A)
		}
	}

	fmt.Println()
}

func (g *Game) printRow() {
	for i := 0; i < WORD_LENGTH; i++ {
		fmt.Print("[ ]")
		// TODO:
		// [ ] not found
		// [?] wrong location
		// [*] found
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

		if len(guess) == WORD_LENGTH {
			break
		} else {
			fmt.Printf("Invalid length. Please enter %v letters\n", WORD_LENGTH)
		}
	}

	g.currentRound().setGuess(guess)

	return nil
}
