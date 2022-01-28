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
		name       string
		word       string
		letter     string
		position   int
		green_eval []int // indicates which letters were marked green
		want       bool
	}{
		{
			name:       "Yellow",
			word:       "FROGS",
			letter:     "F",
			position:   1,
			green_eval: []int{BLACK, BLACK, BLACK, BLACK, BLACK},
			want:       true,
		},
		{
			name:       "Green (Yellow is false) ",
			word:       "FROGS",
			letter:     "F",
			position:   0,
			green_eval: []int{GREEN, BLACK, BLACK, BLACK, BLACK},
			want:       false,
		},
		{
			name:       "Not in word",
			word:       "FROGS",
			letter:     "Z",
			position:   2,
			green_eval: []int{BLACK, BLACK, BLACK, BLACK, BLACK},
			want:       false,
		},
		{
			name:       "Out of bounds (under)",
			word:       "FROGS",
			letter:     "F",
			position:   -1,
			green_eval: []int{BLACK, BLACK, BLACK, BLACK, BLACK},
			want:       false,
		},
		{
			name:       "Out of bounds (over)",
			word:       "FROGS",
			letter:     "F",
			position:   5,
			green_eval: []int{BLACK, BLACK, BLACK, BLACK, BLACK},
			want:       false,
		},
		{
			name:   "double letter (green-yellow)",
			word:   "LOOPY",
			letter: "O", position: 2,
			green_eval: []int{BLACK, GREEN, BLACK, BLACK, BLACK},
			want:       true,
		},
		{
			name:   "double letter (yellow-green)",
			word:   "LOOPY",
			letter: "O", position: 1,
			green_eval: []int{BLACK, BLACK, GREEN, BLACK, BLACK},
			want:       true,
		},
		{
			name:   "double letter (yellow-black, first letter)",
			word:   "LOOPY",
			letter: "O", position: 1,
			green_eval: []int{BLACK, BLACK, BLACK, BLACK, BLACK},
			want:       true,
		},
		{
			name:   "double letter (yellow-black, second letter)",
			word:   "LOOPY",
			letter: "O", position: 2,
			green_eval: []int{BLACK, BLACK, BLACK, BLACK, BLACK},
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWord(tt.word)
			got := w.isYellow(tt.letter, tt.position, tt.green_eval)
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
		letter string
		want   []int
	}{
		{
			name:   "no occurrences",
			word:   "ABCDE",
			letter: "Z",
			want:   []int{0, 1, 2, 3, 4},
		},
		{
			name:   "1 occurrence",
			word:   "ABCDE",
			letter: "A",
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "2 occurrences",
			word:   "ALOHA",
			letter: "A",
			want:   []int{1, 2, 3},
		},
		{
			name:   "3 occurrences",
			word:   "MOMMA",
			letter: "M",
			want:   []int{1, 4},
		},
		{
			name:   "4 occurrences",
			word:   "ABAAA",
			letter: "A",
			want:   []int{1},
		},
		{
			name:   "5 occurrences",
			word:   "AAAAA",
			letter: "A",
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := yellowIndices(tt.word, tt.letter)
			if len(got) != len(tt.want) {
				t.Errorf("yellowIndices(%s,%s) got %v, want %v", tt.word, tt.letter, got, tt.want)
			}
			// TODO slice compare
		})
	}

}
