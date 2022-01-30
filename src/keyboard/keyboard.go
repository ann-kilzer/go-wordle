package keyboard

import (
	"fmt"
	"strings"
)

const UPPER_A = 65
const LOWER_A = 97

// style
const UPPERCASE = 0
const LOWERCASE = 1
const INVISIBLE = 2

type CharacterPadding struct {
	Left  rune
	Right rune
	style int
}

// Format returns the character formatted with the correct padding and style
func (c *CharacterPadding) Format(letter string) string {
	switch c.style {
	case UPPERCASE:
		return fmt.Sprintf("%c%v%c ", c.Left, strings.ToUpper(letter), c.Right)
	case LOWERCASE:
		return fmt.Sprintf("%c%v%c ", c.Left, strings.ToLower(letter), c.Right)
	case INVISIBLE:
		return fmt.Sprintf("%c  %c ", c.Left, c.Right)
	}
	return ""
}

func NewCharacterPadding(l, r rune) *CharacterPadding {
	return &CharacterPadding{
		Left:  l,
		Right: r,
	}
}

// for letterRecord
const UNUSED = 0
const NO_MATCH = 1
const MATCH = 2

// Keyboard tracks used and found letters, and
type Keyboard struct {
	letterRecord   [26]int // track used and found letters
	matchPadding   *CharacterPadding
	noMatchPadding *CharacterPadding
	unusedPadding  *CharacterPadding
}

// todo: allow custom paddings
func NewKeyboard() *Keyboard {
	return &Keyboard{
		matchPadding:   NewCharacterPadding('[', ']'),
		noMatchPadding: NewCharacterPadding(' ', ' '),
		unusedPadding:  NewCharacterPadding('<', '>'),
	}
}

func (k *Keyboard) MarkMatch(letter byte) {
	index := letter - UPPER_A
	if k.letterRecord[index] == UNUSED { // already marked
		k.letterRecord[index] = MATCH
	}
}

func (k *Keyboard) MarkNoMatch(letter byte) {
	index := letter - UPPER_A
	if k.letterRecord[index] == UNUSED { // already marked
		k.letterRecord[index] = NO_MATCH
	}
}

func (k *Keyboard) String() string {
	var sb strings.Builder
	qw := QwertyOrder()
	for row := 0; row < len(qw); row++ {
		for col := 0; col < len(qw[row]); col++ {
			letter := string(qw[row][col])
			alphaIndex := qw[row][col] - UPPER_A
			fmt.Printf("%b", alphaIndex)

			switch k.letterRecord[alphaIndex] {
			case UNUSED:
				sb.WriteString(k.unusedPadding.Format(letter))
			case NO_MATCH:
				sb.WriteString(k.noMatchPadding.Format(" "))
			case MATCH:
				sb.WriteString(k.matchPadding.Format(letter))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
