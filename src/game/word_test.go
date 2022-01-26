package game

import "testing"

func TestIsGreen(t *testing.T) {
	var tests = []struct {
		name     string
		word     string
		letter   string
		position int
		want     bool
	}{
		{name: "Green", word: "STEAM", letter: "A", position: 3, want: true},
		{name: "Wrong location", word: "STEAM", letter: "A", position: 2, want: false},
		{name: "Out of bounds (under)", word: "STEAM", letter: "A", position: -1, want: false},
		{name: "Out of bounds (over)", word: "STEAM", letter: "A", position: 5, want: false},
		{name: "Not found", word: "STEAM", letter: "Z", position: 0, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWord(tt.word)
			got := w.isGreen(tt.letter, tt.position)
			if got != tt.want {
				t.Errorf("isGreen(%s) got %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
