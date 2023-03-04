package kata

func Solution(word string) (reversed string) {
	for i := range word {
		reversed = string(word[i]) + reversed
	}
	// Lordi = idroL
	return reversed
}
