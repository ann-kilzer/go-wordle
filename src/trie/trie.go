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

func newRootNode() *Node {
	return &Node{
		letter:  START_CHAR,
		wordEnd: false,
	}
}

type Trie struct {
	root  *Node
	count int
}

// NewTrie creates a new trie
func NewTrie() *Trie {
	return &Trie{
		root: newRootNode(),
	}
}

// Length is the number of words in the Trie
func (t *Trie) Length() int {
	return t.count
}

// Insert adds the word to the trie
func (t *Trie) Insert(word string) {
	if t.root == nil {
		t.root = newRootNode()
	}
	if len(word) == 0 {
		return
	}

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
	// mark the wordEnd and increment count
	if !prev.wordEnd {
		prev.wordEnd = true
		t.count += 1
	}
}

// HasWord checks if the word exists in the Trie
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
