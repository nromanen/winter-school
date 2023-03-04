package kata

func Gcdi(x, y int) int {
	x = abs(x)
	y = abs(y)
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
	x = abs(x)
	y = abs(y)
	return (x * y) / Gcdi(x, y)
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, acc int) []int {
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = f(acc, arr[i])
		acc = res[i]
	}
	return res
}
