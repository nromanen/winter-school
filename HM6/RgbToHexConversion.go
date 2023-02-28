package kata

//https://www.codewars.com/kata/513e08acc600c94f01000001/train/go

import "fmt"

func getValid(x int) int {
	switch {
	case x < 0:
		return 0
	case x > 255:
		return 255
	default:
		return x
	}
}

func RGB(r, g, b int) string {
	result := fmt.Sprintf("%02X%02X%02X", getValid(r), getValid(g), getValid(b))
	fmt.Println(result)
	return result
}
