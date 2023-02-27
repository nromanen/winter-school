package kata

import "fmt"

func ComputeDepth(n uint) uint {
	var digits [10]bool
	var count uint

	for {
		count++
		for _, digit := range fmt.Sprintf("%d", n*count) {
			digits[digit-'0'] = true
		}
		if allDigitsPresent(digits) {
			return count
		}
	}
}
func allDigitsPresent(digits [10]bool) bool {
	for _, present := range digits {
		if !present {
			return false
		}
	}
	return true
}
