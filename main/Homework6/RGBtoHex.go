package kata

import "fmt"


//Task 5 https://www.codewars.com/kata/513e08acc600c94f01000001
func getValid(x int) int {
  switch {
    case x < 0: return 0
    case x > 255: return 255
    default: return x
  }
}

func RGB(r, g, b int) string {
  res := fmt.Sprintf("%02X%02X%02X", getValid(r), getValid(g), getValid(b))
  fmt.Println(res)
  return res
}