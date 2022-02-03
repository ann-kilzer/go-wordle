package answer

import (
	"fmt"
	"strings"

	"github.com/ann-kilzer/go-wordle/common"
)

// Word represents the answer to a Wordle game
type Word struct {
	value string
}

func NewWord(word string) Word {
	return Word{
		value: strings.ToUpper(word),
	}
}

func (w Word) String() string {
	return w.value
}

// EvaluateGuess determines what the game response should be
// based on evaluating the user's guess against the Word
func (w *Word) EvaluateGuess(guess string) ([common.WORD_LENGTH]int, error) {
	var res [common.WORD_LENGTH]int
	if len(guess) < common.WORD_LENGTH {
		return res, fmt.Errorf("Invalid guess of length %d, expected %d", len(guess), common.WORD_LENGTH)
	}

	for i := 0; i < common.WORD_LENGTH; i++ {
		letter := string(guess[i])
		if w.isGreen(letter, i) {
			res[i] = common.GREEN
		} else if w.isYellow(letter, i, guess) {
			res[i] = common.YELLOW
		} else {
			res[i] = common.BLACK
		}
	}

	return res, nil
}

func (w *Word) IsWin(guess string) bool {
	return guess == w.value
}

func validPosition(position int) bool {
	return position >= 0 && position < common.WORD_LENGTH
}

// isGreen means the letter is in the word and in the correct position
func (w *Word) isGreen(letter string, position int) bool {
	return validPosition(position) && string(w.value[position]) == letter
}

// yellowIndices returns all potential
func yellowIndices(word, guess, letter string) []int {
	res := make([]int, 0)
	for i := 0; i < common.WORD_LENGTH; i++ {
		if string(word[i]) != letter && string(guess[i]) == letter {
			res = append(res, i)
		}
	}

	return res
}

// numGreenForLetter returns the number of green squares for the letter
func numGreenForLetter(word, guess, letter string) int {
	num := 0
	for i := 0; i < common.WORD_LENGTH; i++ {
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

	for i := 0; i < budget && i < len(possibleYellow); i++ {
		if possibleYellow[i] == position {
			return true
		}
	}

	return false
}
