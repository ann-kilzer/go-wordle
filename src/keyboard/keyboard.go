package keyboard

import (
	"fmt"
	"strings"
)

const UPPER_A = 65
const LOWER_A = 97

// "color"
const BLACK = 0
const YELLOW = 1
const GREEN = 2

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

// Keyboard tracks used and found letters, and
type Keyboard struct {
	letterRecord  [26]int // track used and found letters
	greenPadding  *CharacterPadding
	yellowPadding *CharacterPadding
	blackPadding  *CharacterPadding
}

// todo: allow custom paddings
func NewKeyboard() *Keyboard {
	return &Keyboard{
		greenPadding:  NewCharacterPadding('[', ']'),
		yellowPadding: NewCharacterPadding('(', ')'),
		blackPadding:  NewCharacterPadding('<', '>'),
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
			case BLACK:
				sb.WriteString(k.blackPadding.Format(letter))
			case YELLOW:
				sb.WriteString(k.yellowPadding.Format(letter))
			case GREEN:
				sb.WriteString(k.greenPadding.Format(letter))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
