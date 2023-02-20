package kata

//https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go

func Solution(word string) string {
	result := ""
	n := len(word)

	for i := range word {
		result = result + string(word[n-i-1])
	}
	return result
}
