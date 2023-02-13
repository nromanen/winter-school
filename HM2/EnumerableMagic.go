package kata

func EachCons(arr []int, n int) [][]int {
	var result [][]int
	for i := 0; i <= len(arr)-n; i++ {
		result = append(result, arr[i:i+n])
	}
	return result
}
