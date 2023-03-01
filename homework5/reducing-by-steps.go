package main

import (
	"fmt"
	"math"
)

// som : (x, y) -> x + y
func som(x, y int) int {
	return x + y
}

// mini : (x, y) -> min(x, y)
func mini(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

// maxi : (x, y) -> max(x, y)
func maxi(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// lcmu : (x, y) -> lcm(abs(x), abs(y))
func lcmu(x, y int) int {
	return int(math.Abs(float64(x)*float64(y)) / float64(gcdi(x, y)))
}

// gcdi : (x, y) -> gcd(abs(x), abs(y))
func gcdi(x, y int) int {
	a, b := int(math.Abs(float64(x))), int(math.Abs(float64(y)))
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// operArray applies the given function to the array and returns the resulting array
func operArray(fct func(int, int) int, arr []int, init int) []int {
	result := make([]int, len(arr))
	result[0] = fct(init, arr[0])
	for i := 1; i < len(arr); i++ {
		result[i] = fct(result[i-1], arr[i])
	}
	return result
}

func main() {
	a := []int{18, 69, -90, -78, 65, 40}
	fmt.Println(operArray(gcdi, a, a[0])) // [18 3 3 3 1 1]
	fmt.Println(operArray(lcmu, a, a[0])) // [18 414 2070 26910 26910 107640]
	fmt.Println(operArray(som, a, 0))     // [18 87 -3 -81 -16 24]
	fmt.Println(operArray(mini, a, a[0])) // [18 18 -90 -90 -90 -90]
	fmt.Println(operArray(maxi, a, a[0])) // [18 69 69 69 69 69]
}
