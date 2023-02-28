package main

import "fmt"

func reverseString(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

func main() {

	fmt.Println(reverseString("world")) // prints "dlrow"
	fmt.Println(reverseString("word"))  // prints "drow"

}
