package game

import (
	"strings"
)

const BLACK = 0
const YELLOW = 1
const GREEN = 2

// Word represents the answer to a Wordle game
type Word struct {
	value string
}

// evaluateGuess determines what the game response should be
// based on evaluating the user's guess against the Word
func (w *Word) evaluateGuess(guess string) []int {
	res := make([]int, WORD_LENGTH)
	if len(guess) < WORD_LENGTH {
		return res
	}

	for i := 0; i < WORD_LENGTH; i++ {
		letter := string(guess[i])
		if w.isGreen(letter, i) {
			res[i] = GREEN
		} else if w.isYellow(letter, i) {
			res[i] = YELLOW
		} else {
			res[i] = BLACK
		}

	}
	return res
}

func (w *Word) isWin(guess string) bool {
	return guess == w.value
}

// isGreen means the letter is in the word and in the correct position
func (w *Word) isGreen(letter string, position int) bool {
	return string(w.value[position]) == letter
}

// isYellow means the letter is in the word and in the incorrect position
// TODO this implementation is wrong
func (w *Word) isYellow(letter string, position int) bool {
	return strings.Contains(w.value, letter)
}
