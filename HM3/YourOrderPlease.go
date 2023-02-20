package kata

import (
	"strconv"
	"strings"
)

//https://www.codewars.com/kata/55c45be3b2079eccff00010f/train/go

func Order(sentence string) string {
	words := strings.Fields(sentence)
	for j := 1; j <= len(words); j++ {
		for i := j - 1; i < len(words); i++ {
			if strings.Contains(words[i], strconv.Itoa(j)) {
				words[i], words[j-1] = words[j-1], words[i]
			}
		}
	}
	return strings.Join(words, " ")
}
