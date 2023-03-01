package kata

import "sort"
import "strings"

func InArray(array1 []string, array2 []string) []string {
    result := make([]string, 0)
    for _, word1 := range array1 {
        for _, word2 := range array2 {
            if strings.Contains(word2, word1) {
                result = append(result, word1)
                break
            }
        }
    }

    sort.Strings(result)

    return result
}