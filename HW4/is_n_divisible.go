// https://www.codewars.com/kata/5545f109004975ea66000086/train/go

package kata

func IsDivisible(n, x, y int) bool {
	if n%x == 0 && n%y == 0 {
		return true
	}
	return false
}
