package trie_test

import (
	"testing"

	"github.com/ann-kilzer/go-wordle/trie"
)

func initTrie(initWords []string) *trie.Trie {
	tr := trie.NewTrie()
	for i := 0; i < len(initWords); i++ {
		tr.Insert(initWords[i])
	}
	return tr
}

func TestInsert(t *testing.T) {
	var tests = []struct {
		name          string
		trie          *trie.Trie
		word          string
		expectedCount int
	}{
		{
			name:          "non-function init",
			trie:          &trie.Trie{},
			word:          "hello",
			expectedCount: 1,
		},
		{
			name:          "NewTrie init",
			trie:          trie.NewTrie(),
			word:          "hello",
			expectedCount: 1,
		},
		{
			name:          "Invalid word",
			trie:          trie.NewTrie(),
			word:          "",
			expectedCount: 0,
		},
		{
			name:          "Invalid word",
			trie:          initTrie([]string{"A", "AA", "AARDVARK"}),
			word:          "ALAMO",
			expectedCount: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.trie.Insert(tt.word)
			count := tt.trie.Length()
			if count != tt.expectedCount {
				t.Errorf("Insert(...) count %v, want %v", count, tt.expectedCount)
			}
		})
	}
}

func TestHasWord(t *testing.T) {
	var tests = []struct {
		name      string
		initWords []string
		word      string
		want      bool
	}{
		{
			name:      "Empty Trie",
			initWords: []string{},
			word:      "frog",
			want:      false,
		},
		{
			name:      "No match",
			initWords: []string{"bacon", "hello", "teeth"},
			word:      "frog",
			want:      false,
		},
		{
			name:      "Case insensitive match",
			initWords: []string{"FROG"},
			word:      "frog",
			want:      true,
		},
		{
			name:      "Match",
			initWords: []string{"ALOHA", "BREAD", "CORNY"},
			word:      "ALOHA",
			want:      true,
		},
		{
			name:      "Partial overlap but no match",
			initWords: []string{"SUPER"},
			word:      "SUPERFLUOUS",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := initTrie(tt.initWords)
			got := tr.HasWord(tt.word)
			if got != tt.want {
				t.Errorf("HasWord(%s) got %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
