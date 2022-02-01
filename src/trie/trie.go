package trie

import (
	"strings"
)

const UPPER_A = 65

const ALPHABET_LEN = 26
const START_CHAR = '*'

// Node is the building block of a Trie
type Node struct {
	letter  byte
	wordEnd bool
	next    [ALPHABET_LEN]*Node
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			wordEnd: false,
		},
	}
}

// Insert adds the word to the trie

func (t *Trie) Insert(word string) {
	w := strings.ToUpper(word)
	var prev *Node
	cur := t.root
	for i := 0; i < len(w); i++ {
		letter := w[i]
		index := letter - UPPER_A
		if cur.next[index] == nil {
			cur.next[index] = &Node{letter: letter}
		}
		prev = cur
		cur = cur.next[index]
	}
	prev.wordEnd = true
}

func (t *Trie) HasWord(word string) bool {
	w := strings.ToUpper(word)
	var prev *Node
	cur := t.root
	for i := 0; i < len(w); i++ {
		letter := w[i]
		index := letter - UPPER_A
		if cur.next[index] == nil {
			return false
		}
		prev = cur
		cur = cur.next[index]
	}

	return prev != nil && prev.wordEnd 
}
