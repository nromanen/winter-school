package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func main() {
	fmt.Println(Order("is2 Thi1s T4est 3a"))
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2"))
}
