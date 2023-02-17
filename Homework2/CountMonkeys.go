package main

import "fmt"

func countMonkeys (n int) []int{
	var s[] int
	for j := 1; j <= n; j++ {
		s = append(s, j)
	}
	return s 
} 

func main() {	
	fmt.Println(countMonkeys(10))
	fmt.Println(countMonkeys(1))
	fmt.Println(countMonkeys(0))


}
