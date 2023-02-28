package kata


func EachCons(arr []int, n int) [][]int {
    result := [][]int{}

    for i := 0; i < len(arr)-n+1; i++ {
        result = append(result, arr[i:i+n])
    }

    return result
}