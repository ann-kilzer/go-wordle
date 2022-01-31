package trie

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
			letter: START_CHAR,
			wordEnd: false,
		},
	}
}