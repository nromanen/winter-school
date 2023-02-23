package main

import "fmt"

func ReversedString(str string) (rstr string) {
	for i := len(str) - 1; i >= 0; i-- {
		rstr += string(str[i])
	}
	return
}

func main() {
	fmt.Println(ReversedString("world"))
	fmt.Println(ReversedString("word"))
}
