package main

import "fmt"

func computeDepth(n int) int {
	digits := make(map[int]bool)
	i := 1
	for {
		for _, c := range fmt.Sprintf("%d", n*i) {
			digits[int(c-'0')] = true
		}
		if len(digits) == 10 {
			return i
		}
		i++
	}
}

func main() {
	n := 42
	depth := computeDepth(n)
	fmt.Printf("The depth of %d is %d\n", n, depth)
}
