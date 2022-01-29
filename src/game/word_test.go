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

func TestIsYellow(t *testing.T) {
	var tests = []struct {
		name     string
		word     string
		letter   string
		position int
		guess    string
		want     bool
	}{
		{
			name:     "Yellow",
			word:     "FROGS",
			letter:   "F",
			position: 1,
			guess:    "AFROS",
			want:     true,
		},
		{
			name:     "Green (Yellow is false) ",
			word:     "FROGS",
			letter:   "F",
			position: 0,
			guess:    "AFROS",
			want:     false,
		},
		{
			name:     "Not in word",
			word:     "FROGS",
			letter:   "Z",
			position: 2,
			guess:    "PIZZA",
			want:     false,
		},
		{
			name:     "Out of bounds (under)",
			word:     "FROGS",
			letter:   "F",
			position: -1,
			guess:    "ZIGGY",
			want:     false,
		},
		{
			name:     "Out of bounds (over)",
			word:     "FROGS",
			letter:   "F",
			position: 5,
			guess:    "TEENY",
			want:     false,
		},
		{
			name:   "double letter (green-yellow)",
			word:   "LOOPY",
			letter: "O", position: 3,
			guess: "BROOM",
			want:  true,
		},
		{
			name:   "double letter (yellow-green)",
			word:   "BROOM",
			letter: "O", position: 0,
			guess: "LOOPY",
			want:  true,
		},
		{
			name:   "TODO",
			word:   "LOOPY",
			letter: "O", position: 1,
			guess: "AFROS",
			want:  true,
		},
		{
			name:   "TODO",
			word:   "LOOPY",
			letter: "O", position: 0,
			guess: "AFROS",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWord(tt.word)
			got := w.isYellow(tt.letter, tt.position, tt.guess)
			if got != tt.want {
				t.Errorf("isYellow(%s) got %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestYellowIndices(t *testing.T) {
	var tests = []struct {
		name   string
		word   string
		guess  string
		letter string
		want   []int
	}{
		{
			name:   "no occurrences",
			word:   "ABCDE",
			guess:  "ZZZZZ",
			letter: "Z",
			want:   []int{0, 1, 2, 3, 4},
		},
		{
			name:   "1 occurrence",
			word:   "ABCDE",
			guess:  "CRABS",
			letter: "A",
			want:   []int{2},
		},
		{
			name:   "2 occurrences",
			word:   "ALOHA",
			guess:  "BANAL",
			letter: "A",
			want:   []int{1, 3},
		},
		{
			name:   "3 occurrences",
			word:   "AZTEC",
			guess:  "BAZAA",
			letter: "A",
			want:   []int{1, 3, 4},
		},
		{
			name:   "4 occurrences",
			word:   "ABAAA",
			guess:  "BABBB",
			letter: "B",
			want:   []int{0, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := yellowIndices(tt.word, tt.guess, tt.letter)
			if len(got) != len(tt.want) {
				t.Errorf("yellowIndices(%s,%s) got %v, want %v", tt.word, tt.letter, got, tt.want)
			}
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("yellowIndices(%s,%s) got %v, want %v", tt.word, tt.letter, got, tt.want)
				}
			}
		})
	}

}
