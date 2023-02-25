// https://www.codewars.com/kata/59b401e24f98a813f9000026/train/go

package kata

func ComputeDepth(n uint) uint {
	var digits [10]bool
	var count, multiple uint

	for {
		count++
		multiple = n * count

		// Iterate through each digit in the multiple
		for multiple > 0 {
			digit := multiple % 10
			multiple /= 10
			digits[digit] = true
		}

		// Check if all digits have been seen
		if allDigitsSeen(digits) {
			return count
		}
	}
}

// Helper function to check if all digits have been seen
func allDigitsSeen(digits [10]bool) bool {
	for _, seen := range digits {
		if !seen {
			return false
		}
	}
	return true
}
