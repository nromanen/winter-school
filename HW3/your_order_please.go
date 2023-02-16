// https://www.codewars.com/kata/55c45be3b2079eccff00010f

package kata

import (
	"strconv"
	"strings"
	"unicode"
)

func Order(sentence string) string {
	if len(sentence) == 0 {
		return ""
	}

	splitWords := strings.Split(sentence, " ")
	output := make([]string, len(splitWords))
	for _, word := range splitWords {
		numIndex := getNumberIndex(word)
		output[numIndex-1] = word
	}

	return strings.Join(output, " ")
}

func getNumberIndex(word string) int {
	for _, char := range word {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			return num
		}
	}
	return -1
}
