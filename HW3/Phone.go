package kata

import "fmt"

func CreatePhoneNumber(numbers [10]uint) (result string) {
	for _, value := range numbers {
		result += fmt.Sprintf("%d", value)
	}
	// 1,2,3,4,5,6,7,8,9,0}) = (123) 456-7890
	return fmt.Sprintf("(%s) %s-%s", result[:3], result[3:6], result[6:])
}
