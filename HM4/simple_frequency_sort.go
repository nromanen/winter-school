import "sort"
//https://www.codewars.com/kata/5a8d2bf60025e9163c0000bc
func SortByFrequencyAndValue(arr []int) []int {
    freq := make(map[int]int)
    for _, v := range arr {
        freq[v]++
    }
    
    pairs := make([][2]int, 0, len(freq))
    for k, v := range freq {
        pairs = append(pairs, [2]int{k, v})
    }
    
    sort.Slice(pairs, func(i, j int) bool {
        if pairs[i][1] == pairs[j][1] {
            return pairs[i][0] < pairs[j][0]
        }
        return pairs[i][1] > pairs[j][1]
    })
    
    result := make([]int, 0, len(arr))
    for _, p := range pairs {
        for i := 0; i < p[1]; i++ {
            result = append(result, p[0])
        }
    }
    return result
}
