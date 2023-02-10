package kata

//Task 1 https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad/train/go
func Invert(arr []int) (res []int) {
	for _, v := range arr {
		res = append(res, v*-1)
	}
	return res
}
