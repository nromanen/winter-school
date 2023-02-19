package main

import (
	"fmt"
	"strings"
)

func Repeat(number int, text string) string {
	return strings.Repeat(text, number)
}

func main() {
	fmt.Print("Input number: ")
	var x int
	fmt.Scan(&x)
	fmt.Print("Input some text: ")
	var text string
	fmt.Scan(&text)
	fmt.Println(Repeat(5, text))
}
