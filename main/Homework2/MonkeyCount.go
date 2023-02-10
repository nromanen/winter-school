package kata

//Task 3 https://www.codewars.com/kata/56f69d9f9400f508fb000ba7/train/go
func monkeyCount(n int) []int {
  a := []int{} 
  
  for i := 1; i <= n; i++ {
    a = append(a, i)
  }
  
  return a
}