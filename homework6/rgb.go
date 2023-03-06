package main

import "fmt"

func RGB(r, g, b int) string {

	r = max(0, min(255, r))
	g = max(0, min(255, g))
	b = max(0, min(255, b))

	hexStr := fmt.Sprintf("%02X%02X%02X", r, g, b)

	return hexStr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	fmt.Println(RGB(255, 255, 255))
	fmt.Println(RGB(255, 255, 300))
	fmt.Println(RGB(0, 0, 0))
	fmt.Println(RGB(148, 0, 211))

}
