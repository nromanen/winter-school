package kata

import "fmt"

func ComputeDepth(n uint) uint {
	var digits [10]bool
	var t uint

	for {
		t++
		for _, digit := range fmt.Sprintf("%d", n*t) {
			digits[digit-'0'] = true
		}
		if allDigits(digits) {
			return t
		}
	}
}
func allDigits(digits [10]bool) bool {
	for _, present := range digits {
		if !present {
			return false
		}
	}
	return true
}
