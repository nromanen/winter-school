package main

import "fmt"

func getDigits(n int) []int {
	var result []int
	k := 10
	for true {
		digit := n % k
		if n <= 0 {
			break
		}
		result = append(result, digit)
	}
	return result
}

/*func ComputeDepth(n uint) uint {
	// your code here
}*/

func main() {
	fmt.Println(getDigits(1488))
}
