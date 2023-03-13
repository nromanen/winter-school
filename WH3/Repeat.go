package main

import (
	"fmt"
	"strings"
)

func Repeat(n int, s string) string {
	return strings.Repeat(s, n)
}

func main() {
	fmt.Println(Repeat(6, "I"))     // "IIIIII"
	fmt.Println(Repeat(5, "Hello")) // "HelloHelloHelloHelloHello"
}
