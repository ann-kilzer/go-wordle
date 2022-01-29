package dictionary

// V0: use a map
// TODO; Update with a Trie to save memory

type ValidGuesses struct {
	words map[string]bool
}

func LoadValidGuesses() (*ValidGuesses, error) {
	wlist, err := readWords("data/wordlist.txt")
	if err != nil {
		return nil, err
	}

	v := &ValidGuesses{
		words: make(map[string]bool, len(wlist)),
	}

	for _, w := range wlist {
		v.words[w] = true
	}

	return v, nil
}

func (v *ValidGuesses) isValid(guess string) bool {
	return v.words[guess]
}
