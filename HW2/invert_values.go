// https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad/go

package kata

func Invert(arr []int) []int {
	var result []int
	for _, num := range arr {
		result = append(result, -num)
	}
	return result
}
