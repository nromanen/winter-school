package main

import (
	"fmt"
)

func reversed(word string) (reversed string) {
	for i := len(word) - 1; i >= 0; i -= 1 {
		reversed += string(word[i])
	}

	return reversed
}

func main() {

	fmt.Println("Task 4 :", reversed("world"))

}
