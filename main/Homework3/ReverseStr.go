package kata

//Task 1 https://www.codewars.com/kata/5168bb5dfe9a00b126000018
func Solution(word string) (reversed string) {
  
  for i := len(word)-1; i >= 0; i-- {
    reversed += word[i:i+1]
  }
  
  return
}