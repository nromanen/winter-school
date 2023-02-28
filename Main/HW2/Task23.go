package kata

func monkeyCount(n int) []int {
    result := make([]int, n)
    for i := 1; i <= n; i++ {
        result[i-1] = i
    }
    return result
}