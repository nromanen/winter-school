package kata

import "math"

func SumCubes(n int) int {
	// your code here
	result := 0
	for i := 1; i <= n; i++ {
		result += int(math.Pow(float64(i), 3))
	}
	return result
}
