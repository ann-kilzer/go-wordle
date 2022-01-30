package keyboard

import (
	"strings"
)

const UPPER_A = 65
const UPPER_Z = 96
const LOWER_A = 97

// style
const UPPERCASE = 0
const LOWERCASE = 1
const INVISIBLE = 2

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
	if letter < UPPER_A || letter > UPPER_Z {
		return
	}
	index := letter - UPPER_A
	if k.letterRecord[index] == UNUSED {
		k.letterRecord[index] = MATCH
	}
}

func (k *Keyboard) MarkNoMatch(letter byte) {
	if letter < UPPER_A || letter > UPPER_Z {
		return
	}
	index := letter - UPPER_A
	if k.letterRecord[index] == UNUSED {
		k.letterRecord[index] = NO_MATCH
	}
}

// String implements the string interface
func (k Keyboard) String() string {
	var sb strings.Builder
	qw := QwertyOrder()
	for row := 0; row < len(qw); row++ {
		for col := 0; col < len(qw[row]); col++ {
			letter := string(qw[row][col])
			alphaIndex := qw[row][col] - UPPER_A

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
