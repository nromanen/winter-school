package main

import "fmt"

func RGB(r, g, b int) string {
	if r > 255 {
		r = 255
	}
	if g > 255 {
		g = 255
	}
	if b > 255 {
		b = 255
	}
	R := fmt.Sprintf("%02X", r)
	G := fmt.Sprintf("%02X", g)
	B := fmt.Sprintf("%02X", b)
	return R + G + B
}

func main() {
	// expected - 000000
	fmt.Println(RGB(0, 0, 0))
	// expected - FFFFFF
	fmt.Println(RGB(300, 300, 300))
	// expected - 2DA361
	fmt.Println(RGB(45, 163, 97))
}
