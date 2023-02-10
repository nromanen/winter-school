package kata

//Task 2 https://www.codewars.com/kata/55fd2d567d94ac3bc9000064/train/go
func RowSumOddNumbers(n int) int {
  res := 0
  for i := (n - 1) * n + 1; i < n * (n + 1); i = i + 2 {
    res = res + i
  }
  return res
}