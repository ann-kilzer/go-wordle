package dictionary

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Dictionary struct {
	words []string
}

func (d *Dictionary) RandomWord() string {
	idx := rand.Intn(len(d.words))
	return d.words[idx]
}

func Load() (*Dictionary, error) {
	rand.Seed(time.Now().UnixNano())
	d := &Dictionary{}

	file, err := os.Open("data/historic.txt")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d.words = append(d.words, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return d, nil
}
