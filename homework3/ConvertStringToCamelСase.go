package main

import (
	"fmt"
	"strings"
)

func CamelCase(str string) string {
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, "_", " ")

	str = strings.Title(str)
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func main() {
	fmt.Println(CamelCase("river-Street-Street-left-mountain-Yellow-desert"))
	fmt.Println(CamelCase("Blue_Samurai_bridge"))
}
