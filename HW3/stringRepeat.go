package main

import (
	"fmt"
	"strings"
)

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func main() {
	fmt.Println(RepeatStr(5, "I"))
	fmt.Println(RepeatStr(5, "KekW"))
}
