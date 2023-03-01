import (
    "sort"
    "strings"
)
//https://www.codewars.com/kata/550554fd08b86f84fe000a58
func InArray(a1 []string, a2 []string) []string {
    set := make(map[string]bool)
    
    for _, s2 := range a2 {
        for _, s1 := range a1 {
            if strings.Contains(s2, s1) {
                set[s1] = true
            }
        }
    }
    
    result := make([]string, 0, len(set))
    for s := range set {
        result = append(result, s)
    }
    sort.Strings(result)
    
    return result
}
