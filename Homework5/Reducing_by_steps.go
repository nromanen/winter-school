package main

import (
	"fmt"
	"math"
)

func som(x, y int) int {
	return x + y
}

func mini(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func maxi(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func gcdi(x, y int) int {
	x = int(math.Abs(float64(x)))
	y = int(math.Abs(float64(y)))
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcmu(x, y int) int {
	return int(math.Abs(float64(x*y))) / gcdi(x, y)
}

func oper_array(f func(int, int) int, arr []int, init int) []int {
	result := make([]int, len(arr))
	result[0] = f(init, arr[0])
	for i := 1; i < len(arr); i++ {
		result[i] = f(result[i-1], arr[i])
	}
	return result
}

func main() {
	a := []int{18, 69, -90, -78, 65, 40}
	fmt.Println(oper_array(gcdi, a, a[0]))
	fmt.Println(oper_array(lcmu, a, a[0]))
	fmt.Println(oper_array(som, a, 0))
	fmt.Println(oper_array(mini, a, a[0]))
	fmt.Println(oper_array(maxi, a, a[0]))
}
