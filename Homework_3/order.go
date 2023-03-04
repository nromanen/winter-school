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
	wList := strings.Fields(words)
	sList := make([]string, len(wList))
	for _, word := range wList {
		pos, _ := strconv.Atoi(strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsNumber(r)
		}))
		sList[pos-1] = word
	}
	return strings.Join(sList, " ")
}

func main() {
	fmt.Println("Task 2 :")
	fmt.Println(order("is2 Thi1s T4est 3a"))
	fmt.Println(order("4of Fo1r pe6ople g3ood th5e the2"))
}
