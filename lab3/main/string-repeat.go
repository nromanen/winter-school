package main

import (
	"fmt"
	"strings"
)

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func main() {
	// expected - "IIIII"
	fmt.Println(RepeatStr(5, "I"))
	// expected - "HelloHelloHelloHelloHello"
	fmt.Println(RepeatStr(5, "Hello"))
}
