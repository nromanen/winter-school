package kata

import "math"
import "fmt"

func RGB(r, g, b int) string {
    r = int(math.Max(0, math.Min(255, float64(r))))
    g = int(math.Max(0, math.Min(255, float64(g))))
    b = int(math.Max(0, math.Min(255, float64(b))))

    return fmt.Sprintf("%02X%02X%02X", r, g, b)
}
