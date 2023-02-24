package kata

//Task 1 https://www.codewars.com/kata/56efab15740d301ab40002ee/train/go

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
func Gcdi(x, y int) int {
	x, y = abs(x), abs(y)
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
	return abs(x) * abs(y) / Gcdi(x, y)
}

type FParam func(int, int) int

func oper(f FParam, x, y int) int {
	return f(x, y)
}
func OperArray(f FParam, arr []int, init int) []int {
	var result []int
	p := init
	for i := 0; i < len(arr); i++ {
		p = oper(f, p, arr[i])
		result = append(result, p)
	}
	return result
}
