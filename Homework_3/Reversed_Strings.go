package kata

func Solution(word string) (result string) {
	for _, char := range word {
		result = string(char) + result
	}
	return
}
