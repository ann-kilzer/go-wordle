package keyboard

import (
	"testing"
)

func TestString(t *testing.T) {
	// simulate 1 round: word: CARAT, guess: ALOHA
	round1kb := *NewKeyboard()
	round1kb.MarkMatch('A')
	round1kb.MarkNoMatch('L')
	round1kb.MarkNoMatch('O')
	round1kb.MarkNoMatch('H')
	round1kb.MarkNoMatch('A') // should not clear match

	// simulate 2 rounds: word: CARAT guesses: ALOHA, STEAM
	round2kb := *NewKeyboard()
	round2kb.MarkMatch('A')
	round2kb.MarkNoMatch('L')
	round2kb.MarkNoMatch('O')
	round2kb.MarkNoMatch('H')
	round2kb.MarkNoMatch('A') // should not clear match

	round2kb.MarkNoMatch('S')
	round2kb.MarkMatch('T')
	round2kb.MarkNoMatch('E')
	round2kb.MarkMatch('A')
	round2kb.MarkNoMatch('M')

	var tests = []struct {
		name     string
		keyboard Keyboard
		want     string
	}{
		{
			name:     "Default Keyboard",
			keyboard: *NewKeyboard(),
			want:     "<Q> <W> <E> <R> <T> <Y> <U> <I> <O> <P> \n<A> <S> <D> <F> <G> <H> <J> <K> <L> \n<Z> <X> <C> <V> <B> <N> <M> \n",
		},
		{
			name:     "Keyboard after first round",
			keyboard: round1kb,
			want:     "<Q> <W> <E> <R> <T> <Y> <U> <I>     <P> \n[A] <S> <D> <F> <G>     <J> <K>     \n<Z> <X> <C> <V> <B> <N> <M> \n",
		},
		{
			name:     "Keyboard after second round",
			keyboard: round2kb,
			want:     "<Q> <W>     <R> [T] <Y> <U> <I>     <P> \n[A]     <D> <F> <G>     <J> <K>     \n<Z> <X> <C> <V> <B> <N>     \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.keyboard.String()
			if got != tt.want {
				t.Errorf("String() got %v, want %v", got, tt.want)
			}
		})
	}
}
