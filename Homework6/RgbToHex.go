package main

import (
	"fmt"
	"math"
)

func rgb(r, g, b int) string {
	r = int(math.Max(0, math.Min(255, float64(r))))
	g = int(math.Max(0, math.Min(255, float64(g))))
	b = int(math.Max(0, math.Min(255, float64(b))))

	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}

func main() {
	fmt.Println(rgb(255, 255, 255)) // FFFFFF
	fmt.Println(rgb(255, 255, 300)) // FFFFFF
	fmt.Println(rgb(0, 0, 0))       // 000000
	fmt.Println(rgb(148, 0, 211))   // 9400D3
	fmt.Println(rgb(223, 223, 223)) // DFDFDF

}
