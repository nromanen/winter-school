package kata

func RowSumOddNumbers(n int) int {
	var result int
	for i := 0; i < n; i++ {
		result += (n * n)
	}
	return result
	//return n*n*n
}
