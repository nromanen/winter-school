package kata

func EachCons(arr []int, n int) (result [][]int) {
	for i := 0; i+n <= len(arr); i++ {
		result = append(result, arr[i:i+n])
	}
	return
}
