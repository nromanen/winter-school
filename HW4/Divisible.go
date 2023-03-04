package kata

func IsDivisible(n, x, y int) bool {

	if n%x == 0 && n%y == 0 {
		return true
	} else {
		return false
	}
	// 1) n =  12, x = 2, y = 6 =>  true =>  12 is divisible by 2 and 6
	// 2) n =  12, x = 7, y = 5 => false =>  12 is neither divisible by 7 nor 5
}
