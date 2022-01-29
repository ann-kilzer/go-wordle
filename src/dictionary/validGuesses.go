package dictionary

import "strings"

// V0: use a map
// TODO; Update with a Trie to save memory

type ValidGuesses struct {
	words map[string]bool
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
		words: make(map[string]bool, len(wl)+len(vg)),
	}

	for _, w := range wl {
		v.words[strings.ToUpper(w)] = true
	}

	for _, w := range vg {
		v.words[strings.ToUpper(w)] = true
	}

	return v, nil
}

func (v *ValidGuesses) Contains(guess string) bool {
	return v.words[guess]
}
