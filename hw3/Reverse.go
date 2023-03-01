package main

import "fmt"

func Reverse(s string) string {
	text := []rune(s)
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		text[i], text[j] = text[j], text[i]
	}
	return string(text)
}

func main() {
	var text string
	fmt.Print("Input some text: ")
	fmt.Scan(&text)

	fmt.Print(Reverse(text))
}

