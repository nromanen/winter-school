package main

import (
	"fmt"
)

func ReverseString(text string) string {
	var textInit = []rune(text)
	// Ok fmt.Println(len(textInit))
	// Ok fmt.Println(textInit[2]) -- ASCII OUTPUT 114 = r
	var textReversed = ""
	for i := len(textInit) - 1; i > -1; i-- {
		letter := textInit[i]
		textReversed = textReversed + string(letter)
	}
	return textReversed
}

func main() {
	fmt.Println(ReverseString("world"))
	fmt.Println(ReverseString("word"))
	fmt.Println(ReverseString("abracadabra"))


}

// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
