package dictionary

import (
	"strings"

	"github.com/ann-kilzer/go-wordle/trie"
)

type ValidGuesses struct {
	words *trie.Trie
}

func LoadValidGuesses() (*ValidGuesses, error) {
	wl, err := readWords("data/wordlist.txt")
	if err != nil {
		return nil, err
	}
	vg, err := readWords("data/validGuesses.txt")
	if err != nil {
		return nil, err
	}

	v := &ValidGuesses{
		words: trie.NewTrie(),
	}

	for _, w := range wl {
		v.words.Insert(strings.ToUpper(w))
	}

	for _, w := range vg {
		v.words.Insert(strings.ToUpper(w))
	}

	return v, nil
}

func (v *ValidGuesses) Contains(guess string) bool {
	return v.words.HasWord(guess)
}
