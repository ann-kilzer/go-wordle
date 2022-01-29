package keyboard

// QuertyOrder returns a 2D byte array in the order of a US QWERTY Keyboard
func QwertyOrder() [][]byte {
	return [][]byte{
		{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'},
		{'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L'},
		{'Z', 'X', 'C', 'V', 'B', 'N', 'M'},
	}
}
