package keyboard

import (
	"fmt"
	"strings"
)

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
