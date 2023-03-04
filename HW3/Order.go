package main

import "fmt"

func EachCons(arr []int, n int) (result [][]int) {
	for i := 0; i < len(arr)-n+1; i++ {
		result = append(result, arr[i:n+i])
	}
	return
}
func main() {

	fmt.Println("for example, ([1,2,3,4], 2)  -", EachCons([]int{1, 2, 3, 4}, 2))

}
