package game

import "github.com/ann-kilzer/go-wordle/common"

// evaluateRound must be called after the guess is read
func (g *Game) evaluateRound() error {
	// evaluate the guess
	eval, err := g.word.EvaluateGuess(g.currentRound().Guess)
	if err != nil {
		return err
	}

	g.currentRound().SetEval(eval)
	return nil
}

// markUsedLetters updates which letters have been used on the letterRecord "keyboard"
func (g *Game) markUsedLetters() {
	guess := g.currentRound().Guess
	eval := g.currentRound().Eval
	for i := 0; i < common.WORD_LENGTH; i++ {
		letterByte := guess[i]
		if eval[i] == common.GREEN {
			g.keyboard.MarkMatch(letterByte)
		} else if eval[i] == common.YELLOW {
			g.keyboard.MarkMatch(letterByte)
		} else {
			g.keyboard.MarkNoMatch(letterByte)
		}
	}
}
