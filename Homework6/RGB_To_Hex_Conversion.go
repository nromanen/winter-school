package main

import "fmt"

func rgb(r, g, b int) string {
	// Round values outside of 0-255 range
	if r < 0 {
		r = 0
	} else if r > 255 {
		r = 255
	}
	if g < 0 {
		g = 0
	} else if g > 255 {
		g = 255
	}
	if b < 0 {
		b = 0
	} else if b > 255 {
		b = 255
	}

	// Convert to hexadecimal string and return
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}

func main() {
	fmt.Println(rgb(255, 255, 255)) // FFFFFF
	fmt.Println(rgb(255, 255, 300)) // FFFFFF
	fmt.Println(rgb(0, 0, 0))       // 000000
	fmt.Println(rgb(148, 0, 211))   // 9400D3
}
