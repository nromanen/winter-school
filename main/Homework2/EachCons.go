package kata

//Task 4 https://www.codewars.com/kata/545af3d185166a3dec001190

func EachCons(arr []int, n int) [][]int {
	var result [][]int
	for i := 0; i <= len(arr) - n; i++ {
	  result = append(result, arr[i:i + n])
	}
	return result
  }