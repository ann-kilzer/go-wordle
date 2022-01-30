package keyboard

import "testing"

func TestString(t *testing.T) {
	round1kb := *NewKeyboard()
	round1kb.MarkMatch('A')
	round1kb.MarkNoMatch('L')
	round1kb.MarkNoMatch('O')
	round1kb.MarkNoMatch('H')

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
