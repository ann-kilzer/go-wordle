package game

import (
	"fmt"
	"strings"

	"github.com/ann-kilzer/go-wordle/common"
	"github.com/ann-kilzer/go-wordle/snsOutput"
)

// printLetters shows the current keyboard state
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
		case common.GREEN:
			fmt.Printf("[%v]", letter)
		case common.YELLOW:
			fmt.Printf("?%v?", strings.ToLower(letter))
		case common.BLACK:
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

		if len(guess) != common.WORD_LENGTH {
			fmt.Printf("Invalid length. Please enter %v letters\n", common.WORD_LENGTH)
		} else if g.validGuesses.Contains(guess) {
			break
		} else {
			fmt.Println("Invalid word")
		}
	}

	g.currentRound().SetGuess(guess)

	return nil
}

// PrintSNSOutput prints the SNS formatted grid for sharing
// after a completed game
func (g *Game) PrintSNSOutput() {
	var evals []common.Evaluation
	for r := 0; r < len(g.rounds); r++ {
		if g.rounds[r] == nil {
			break
		}
		evals = append(evals, g.rounds[r].Eval)
	}

	fmt.Print(snsOutput.GenerateOutput(evals))
}
