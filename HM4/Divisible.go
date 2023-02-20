package kata

// https://www.codewars.com/kata/5545f109004975ea66000086/train/go

func IsDivisible(n, x, y int) bool {
	if (n%x)+(n%y) == 0 {
		return true
	}
	return false
}
