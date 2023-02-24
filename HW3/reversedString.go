package main

import "fmt"

func reverse(word string) string {
	result := ""
	n := len(word)
	for i := range word {
		result = result + string(word[n-i-1])
	}
	return result
}

func main() {
	fmt.Println(reverse("world"))
	fmt.Println(reverse("word"))
}
