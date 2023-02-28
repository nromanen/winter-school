package kata

//https://www.codewars.com/kata/56efab15740d301ab40002ee/train/go

import (
	"math"
)

func Gcdi(x, y int) int {
	x, y = int(math.Abs(float64(x))), int(math.Abs(float64(y)))
	for y > 0 {
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
