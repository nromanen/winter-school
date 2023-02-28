package kata

func Invert(array []int) []int {
    result := make([]int, len(array))
    for i, num := range array {
        result[i] = num * -1
    }
    return result
}