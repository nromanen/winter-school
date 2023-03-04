// https://www.codewars.com/kata/56efab15740d301ab40002ee/train/go

package kata

import (
	"math"
)

func Gcdi(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return int(math.Abs(float64(x)))
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
	return int(math.Abs(float64(x*y)) / float64(Gcdi(x, y)))
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	result := make([]int, 0, len(arr))
	acc := init
	for _, i := range arr {
		acc = f(acc, i)
		result = append(result, acc)
	}
	return result
}