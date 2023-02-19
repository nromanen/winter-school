package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func order(words string) string {
	if len(words) == 0 {
		return ""
	}
	wordList := strings.Fields(words)
	sortedList := make([]string, len(wordList))
	for _, word := range wordList {
		pos, _ := strconv.Atoi(strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsNumber(r)
		}))
		sortedList[pos-1] = word
	}
	return strings.Join(sortedList, " ")
}

func main() {
	fmt.Println(order("is2 Thi1s T4est 3a"))
	fmt.Println(order("4of Fo1r pe6ople g3ood th5e the2"))
}
