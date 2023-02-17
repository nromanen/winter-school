package kata

import "fmt"

// Task 2 https://www.codewars.com/kata/525f50e3b73515a6db000b83
func CreatePhoneNumber(numbers [10]uint) (result string) {
	for _, value := range numbers {
		result += fmt.Sprintf("%d", value)
	}

	return fmt.Sprintf("(%s) %s-%s", result[:3], result[3:6], result[6:])
}
