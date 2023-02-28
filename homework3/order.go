package main

import (
	"fmt"
	"strconv"
	"strings"
)

func order(words string) string {
	if words == "" {
		return ""
	}
	wordsSlice := strings.Split(words, " ")
	orderedSlice := make([]string, len(wordsSlice))
	for _, word := range wordsSlice {
		var posStr string
		for _, char := range word {
			if char >= '0' && char <= '9' {
				posStr += string(char)
			}
		}
		pos, _ := strconv.Atoi(posStr)
		orderedSlice[pos-1] = word
	}
	return strings.Join(orderedSlice, " ")
}

func main() {

	fmt.Println(order("is2 Thi1s T4est 3a"))
	fmt.Println(order("4of Fo1r pe6ople g3ood th5e the2"))

}
