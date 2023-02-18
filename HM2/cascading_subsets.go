package kata
//https://www.codewars.com/kata/545af3d185166a3dec001190
func EachCons(arr []int, n int) [][]int {
	var subsets [][]int
	for i := 0; i <= len(arr)-n; i++ {
		subset := arr[i : i+n]
		subsets = append(subsets, subset)
	}
	return subsets
}
