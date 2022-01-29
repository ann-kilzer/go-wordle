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

func NewWord(word string) Word {
	return Word{
		value: word,
	}
}

func (w Word) String() string {
	return w.value
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
		} else {
			res[i] = BLACK
		}
	}

	for i := 0; i < WORD_LENGTH; i++ {
		letter := string(guess[i])
		if w.isYellow(letter, i, guess) {
			res[i] = YELLOW
		}
	}

	return res
}

func (w *Word) isWin(guess string) bool {
	return guess == w.value
}

func validPosition(position int) bool {
	return position >= 0 && position < WORD_LENGTH
}

// isGreen means the letter is in the word and in the correct position
func (w *Word) isGreen(letter string, position int) bool {
	return validPosition(position) && string(w.value[position]) == letter
}

// yellowIndices returns all potential
func yellowIndices(word, guess, letter string) []int {
	res := make([]int, 0)
	for i := 0; i < WORD_LENGTH; i++ {
		if string(word[i]) != letter && string(guess[i]) == letter {
			res = append(res, i)
		}
	}

	return res
}

// numGreenForLetter returns the number of green squares for the letter
func numGreenForLetter(word, guess, letter string) int {
	num := 0
	for i := 0; i < WORD_LENGTH; i++ {
		if string(guess[i]) == letter && string(word[i]) == letter {
			num += 1
		}
	}
	return num
}

// isYellow means the letter is in the word and in the incorrect position
func (w *Word) isYellow(letter string, position int, guess string) bool {
	if !validPosition(position) || string(w.value[position]) == letter {
		return false
	}

	budget := strings.Count(w.value, letter) - numGreenForLetter(w.value, guess, letter)

	possibleYellow := yellowIndices(w.value, guess, letter)
	if len(possibleYellow) == 0 {
		return false
	}

	// todo: calculate the letter budget then loop through the yellows
	for i := 0; i < budget && i < len(possibleYellow); i++ {
		if possibleYellow[i] == position {
			return true
		}
	}

	return false
}
