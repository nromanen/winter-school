// https://www.codewars.com/kata/5168bb5dfe9a00b126000018

package kata

func Solution(word string) string {
	// Convert the string to a slice of runes
	r := []rune(word)

	// Loop through the slice and swap characters from both ends
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
