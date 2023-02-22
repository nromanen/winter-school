package main

import (
	"fmt"
	"math"
)

func Gcdi(x, y int) int {
	x, y = int(math.Abs(float64(x))), int(math.Abs(float64(y)))
	for y > 0 {
		x, y = y, x%y
	}
	return x
}
func Sum(x, y int) int {
	return x + y
}
func Maxi(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func Mini(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func Lcmu(x, y int) int {
	return int(math.Abs(float64(x*y))) / Gcdi(x, y)
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) (r []int) {
	r = append(r, f(init, arr[0]))
	for i := 1; i < len(arr); i++ {
		r = append(r, f(r[i-1], arr[i]))
	}
	return
}

func main() {
	a := []int{18, 69, -90, -78, 65, 40}
	// expected - [18, 87, -3, -81, -16, 24]
	fmt.Println(OperArray(Sum, a, 0))
	// expected - [18, 18, -90, -90, -90, -90]
	fmt.Println(OperArray(Mini, a, a[0]))
	// expected - [18, 69, 69, 69, 69, 69]
	fmt.Println(OperArray(Maxi, a, a[0]))
	// expected - [18, 3, 3, 3, 1, 1]
	fmt.Println(OperArray(Gcdi, a, a[0]))
	// expected - [18, 414, 2070, 26910, 26910, 107640]
	fmt.Println(OperArray(Lcmu, a, a[0]))
}
