package main

import (
	"fmt"
)

func each_cons(s []int, n int) (res [][]int) {
	if n > len(s) {
		fmt.Println("Looks like n is bigger than lenght of array:")
		return
	}
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
		fmt.Println(s[i:i+n])
	}
	return
}

func main() {
	fmt.Println(each_cons([]int{1, 2, 3, 4}, 1))
	fmt.Println(each_cons([]int{1, 2, 3, 4}, 2))
	fmt.Println(each_cons([]int{1, 2, 3, 4}, 3))
	fmt.Println(each_cons([]int{1, 2, 3, 4}, 5))
	fmt.Println(each_cons([]int{1, 2, 3, 4}, 2))

}
