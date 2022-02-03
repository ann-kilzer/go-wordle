package round

import "strings"

const WORD_LENGTH = 5
const EMPTY_GUESS = "     " // 5 blank spaces

type Round struct {
	Guess string           // user input
	Eval  [WORD_LENGTH]int // the evaluation of the guess
}

func NewRound() *Round {
	return &Round{Guess: EMPTY_GUESS}
}

// setGuess records the guess in the round as an uppercase string
func (r *Round) SetGuess(guess string) {
	r.Guess = strings.ToUpper(guess)
}

// setEval records the evaluation in the round
func (r *Round) SetEval(eval [WORD_LENGTH]int) {
	r.Eval = eval
}
