package kata

import (
    "sort"
)

func Solve(arr []int) []int {
    freq := make(map[int]int)
    for _, val := range arr {
        freq[val]++
    }

    sorted := make([]int, len(arr))
    copy(sorted, arr)

    sort.Slice(sorted, func(i, j int) bool {
        if freq[sorted[i]] != freq[sorted[j]] {
            return freq[sorted[i]] > freq[sorted[j]]
        }
        return sorted[i] < sorted[j]
    })

    return sorted
}
