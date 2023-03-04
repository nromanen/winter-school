package main

import "fmt"

func Invert(arr []int) []int {
	reverse := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		reverse[i] = (arr[i] * (-1))
	}
	return reverse
}
func main() {
	fmt.Println("for example: invert[1,2,3,4,5] ==", Invert([]int{1, 2, 3, 4, 5}))

}
