package kata

import "math"

func Gcdi(x, y int) int {
	x, y = int(math.Abs(float64(x))), int(math.Abs(float64(y)))
	for y != 0 {
		x, y = y, x%y
	}
	return x

}
func Som(x, y int) int {
	return x + y
}
func Maxi(x, y int) int {
	if x < y {
		return y
	} else {
		return x
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

func OperArray(f FParam, arr []int, init int) []int {
	result := []int{}

	cop := init

	for _, r := range arr {
		cop = f(cop, r)
		result = append(result, cop)
	}

	return result
}
