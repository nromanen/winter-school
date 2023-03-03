package kata
//https://www.codewars.com/kata/59b401e24f98a813f9000026
func ComputeDepth(n int) int {
    digits := make(map[int]bool)
    count := 0

    for len(digits) < 10 {
        count++
        multiple := n * count

        for multiple > 0 {
            digit := multiple % 10
            digits[digit] = true
            multiple /= 10
        }
    }

    return count
}
