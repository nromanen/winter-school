package kata

func RowSumOddNumbers(n int) int {
	s := 1
	for i := 1; i < n; i++ {
		s += i
	}
	a := 2*s - 1
	result := 0
	for i := 0; i < n; i++ {
		result += a + 2*i
	}
	return result
}
