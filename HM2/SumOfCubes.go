package kata

func SumCubes(n int) int {
	var i, sum int

	sum = 0

	for i = 1; i <= n; i = i + 1 {
		sum = sum + i*i*i
	}
	return sum
}
