package main

import "fmt"

func RGB(r, g, b int) string {
	return Limit(r) + Limit(g) + Limit(b)
}

func Limit(rgb int) string {

	if rgb < 0 {
		rgb = 0
	}

	if rgb > 255 {
		rgb = 255
	}
	return fmt.Sprintf("%02X", rgb)
}

func main() {
	fmt.Println(RGB(0, 0, 0))
	fmt.Println(RGB(216, 11, 3))
	fmt.Println(RGB(81, 255, 47))
}
