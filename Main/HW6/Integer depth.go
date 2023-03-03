package kata
import "fmt"
func ComputeDepth(n uint) uint {
    seen := make(map[rune]bool)
    for i := uint(1); ; i++ {
        for _, c := range fmt.Sprintf("%d", n*i) {
            seen[c] = true
        }
        if len(seen) == 10 {
            return i
        }
    }
}