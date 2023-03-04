package main

import (
	"fmt"
	"strings"
)

func CamelCase(w string) string {
	w = strings.ReplaceAll(w, "-", ",")
	w = strings.ReplaceAll(w, "_", ",")

	arr := strings.Split(w, ",")
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res += strings.Title(arr[i])
	}
	return res

}

func main() {

	w := "You_like-to_code"
	fmt.Println("Task 1 :", CamelCase(w))
}
