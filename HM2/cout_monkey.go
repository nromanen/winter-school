// https://www.codewars.com/kata/55fd2d567d94ac3bc9000064
package main

import "fmt"

func main() {
	m := []int{}
	n := 10
	for i := 1; i < n; i++ {
		m = append(m, i)
	}
	fmt.Print(m)
}
