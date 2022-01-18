package game

import (
	"bytes"
	"fmt"
)

const A = 65

type Round struct {
	guess    []rune // user input
	feedback string // the response to the user
}

func (g *Game) printLetters() {
	// TODO
	for i := 0; i < len(g.usedLetters); i++ {
		if g.usedLetters[i] {
			fmt.Print("x")
		} else {
			fmt.Printf("%c", i+A)
		}
	}

	fmt.Println()
}

func (g *Game) readLetter() (rune, error) {
	char, _, err := g.reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	return char, err
}

func (g *Game) printRow() {
	for i := 0; i < WORD_LENGTH; i++ {
		fmt.Print("[ ]")
		// TODO:
		// [ ] not found
		// [?] wrong location
		// [*] found
	}

	fmt.Println()
}

func (g *Game) readGuess() error {
	fmt.Print(">")

	var b bytes.Buffer

	for i := 0; i < WORD_LENGTH; i++ {
		r, err := g.readLetter()
		if err != nil {
			return err
		}

		b.WriteRune(r)
	}

	fmt.Println(b.String())
	// TODO save guess

	return nil
}
