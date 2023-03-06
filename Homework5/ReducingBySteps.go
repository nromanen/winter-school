package main

import (
	"fmt"
	"math"
)

func Gcdi(x, y int) int {
	x = int(math.Abs(float64(x)))
	y = int(math.Abs(float64(y)))
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func Som(x, y int) int {
	return x + y
}

func Maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Lcmu(x, y int) int {
	return int(math.Abs(float64(x*y))) / Gcdi(x, y)
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	result := make([]int, len(arr))
	for i, val := range arr {
		if i == 0 {
			result[i] = f(init, val)
		} else {
			result[i] = f(result[i-1], val)
		}
	}
	return result

}

func main() {
	a := []int{18, 69, -90, -78, 65, 40}
	fmt.Println(OperArray(Gcdi, a, a[0])) // [18, 3, 3, 3, 1, 1]
	fmt.Println(OperArray(Lcmu, a, a[0])) // [18, 414, 2070, 26910, 26910, 107640]
	fmt.Println(OperArray(Som, a, 0))     // [18, 87, -3, -81, -16, 24]
	fmt.Println(OperArray(Mini, a, a[0])) // [18, 18, -90, -90, -90, -90]
	fmt.Println(OperArray(Maxi, a, a[0])) // [18, 69, 69, 69, 69, 69]

}
