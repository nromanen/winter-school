package main

import "fmt"

func lcm(temp1 int, temp2 int) int {
	var lcmnum int = 1
	if temp1 > temp2 {
		lcmnum = temp1
	} else {
		lcmnum = temp2
	}
	for {
		if lcmnum%temp1 == 0 && lcmnum%temp2 == 0 {
			break
		}
		lcmnum++
	}
	return lcmnum
}

func gcd(temp1 int, temp2 int) int {
	var gcdnum int
	for i := 1; i <= temp1 && i <= temp2; i++ {
		if temp1%i == 0 && temp2%i == 0 {
			gcdnum = i
		}
	}
	return gcdnum
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func Gcdi(x, y int) int {
	return gcd(abs(x), abs(y))
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
	return lcm(abs(x), abs(y))
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			result = append(result, f(init, arr[i]))
		} else {
			result = append(result, f(result[i-1], arr[i]))
		}
	}
	return result
}

func main() {
	a := []int{18, 69, -90, -78, 65, 40}
	fmt.Println(OperArray(Gcdi, a, a[0]))
}
