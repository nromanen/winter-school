package kata

import (
	"fmt"
)

func RGB(r, g, b int) string {
	args := [3]int{r, g, b}
	result := ""

	for _, v := range args {

		if v > 255 {
			v = 255
		} else if v < 0 {
			v = 0
		}
		vs := fmt.Sprintf("%02X", v)
		result = result + vs
	}
	return result
}
