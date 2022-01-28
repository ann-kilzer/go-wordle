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
		}
	}

	for i := 0; i < WORD_LENGTH; i++ {
		letter := string(guess[i])
		if w.isYellow(letter, i, res) {
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

func validPosition(position int) bool {
	return position >= 0 && position < WORD_LENGTH
}

// isGreen means the letter is in the word and in the correct position
func (w *Word) isGreen(letter string, position int) bool {
	return validPosition(position) && string(w.value[position]) == letter
}

// indices returns all indices that contain letter l
func indices(s, l string) []int {
	res := make([]int, 0)
	for i := 0; i < WORD_LENGTH; i++ {
		letter := string(s[i])
		if l == letter {
			res = append(res, i)
			// fmt.Println("✨", i, l)
		}
	}

	return res
}

// isYellow means the letter is in the word and in the incorrect position
// TODO this implementation is wrong
func (w *Word) isYellow(letter string, position int, eval []int) bool {
	if !validPosition(position) || string(w.value[position]) == letter {
		return false
	}

	first_index := strings.Index(w.value, letter)
	if first_index == -1 || eval[first_index] == GREEN {
		return false
	}

	return strings.Contains(w.value, letter)
}
