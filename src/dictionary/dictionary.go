package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Dictionary struct {
	words []string
}

// RandomWord returns a random word from the wordlist dictionary
func (d *Dictionary) RandomWord() string {
	idx := rand.Intn(len(d.words))
	return d.words[idx]
}

func readWords(filename string) ([]string, error) {
	words := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return words, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return words, err
	}

	return words, nil
}

func LoadDictionary() (*Dictionary, error) {
	rand.Seed(time.Now().UnixNano())
	d := &Dictionary{}

	var err error
	d.words, err = readWords("data/wordlist.txt")
	if err != nil {
		return nil, err
	}

	return d, nil
}
