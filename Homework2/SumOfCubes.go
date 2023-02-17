package main

import "fmt"

func SumCubesOfNumber (n int) int{
	var sum int
	for j := 1; j <= n; j++ {
		sum += j*j*j
	}
	return sum 	
	
}

func main() {	
	fmt.Println(SumCubesOfNumber(2))
	fmt.Println(SumCubesOfNumber(3))
	fmt.Println(SumCubesOfNumber(4))


}
