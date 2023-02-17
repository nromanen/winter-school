package kata

//Task 4 https://www.codewars.com/kata/57a0e5c372292dd76d000d7e
func RepeatStr(repetitions int, value string) string {
  var result = ""
    
  for i := 0; i < repetitions; i++ {
    result += value
  }
  return result
}