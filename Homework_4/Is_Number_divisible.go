package kata

func IsDivisible(n, x, y int) bool {
	if (n%x)+(n%y) == 0 {
		return true
	}
	return false
}
