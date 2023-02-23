package main

import (
	"fmt"
	"strings"
)

func StringRepeat(str string, n int) string {
	return strings.Repeat(str, n)
}

func main() {
	fmt.Println(StringRepeat("I", 5))
	fmt.Println(StringRepeat("Hello", 7))
	fmt.Println(StringRepeat("Airport? I wont go to the airport", 3))
	fmt.Println(StringRepeat("CoconutsHaveWaterInThem", 2))
}
