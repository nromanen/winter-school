package kata

import (
	"fmt"
	"math"
)

func RGB(r, g, b int) string {
	r = int(math.Min(math.Max(float64(r), 0), 255) + 0.5)
	g = int(math.Min(math.Max(float64(g), 0), 255) + 0.5)
	b = int(math.Min(math.Max(float64(b), 0), 255) + 0.5)
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}
