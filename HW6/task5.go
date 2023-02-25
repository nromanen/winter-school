// https://www.codewars.com/kata/513e08acc600c94f01000001/train/go

package kata

import "fmt"

func RGB(r, g, b int) string {
	// Round the values to the nearest valid values
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

	// Convert the decimal values to hexadecimal strings and concatenate them
	hex := fmt.Sprintf("%02X%02X%02X", r, g, b)
	return hex
}
