// https://www.codewars.com/kata/56f69d9f9400f508fb000ba7/go

package kata

func monkeyCount(n int) []int {
	var arr []int
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	return arr
}
