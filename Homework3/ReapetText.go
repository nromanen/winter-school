package main

import (
	"fmt"
	)


func Repeat(text string, n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res = res + text
	}
	return res
}

func main() {
	fmt.Println(Repeat("V", 5))
	fmt.Println(Repeat("I", 6))
	fmt.Println(Repeat("Hello", 5))
	
}

//6, "I"     -> "IIIIII"
//5, "Hello" -> "HelloHelloHelloHelloHello"