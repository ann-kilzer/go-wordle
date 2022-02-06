package snsOutput

import (
	"fmt"
	"strings"

	"github.com/ann-kilzer/go-wordle/common"
)

func GenerateOutput(evals []common.Evaluation) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Go-Wordle %d/%d\n", len(evals), common.ROUNDS))

	for r := 0; r < len(evals); r++ {
		e := evals[r]
		for i := 0; i < len(e); i++ {
			sb.WriteRune(colorCodeToRune(e[i]))
		}

		sb.WriteRune('\n')
	}

	return sb.String()
}

func colorCodeToRune(color int) rune {
	switch color {
	case common.BLACK:
		return '⬛'
	case common.YELLOW:
		return '🟨'
	case common.GREEN:
		return '🟩'
	}
	return '⬛'
}
