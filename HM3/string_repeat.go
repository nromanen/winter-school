package kata
//https://www.codewars.com/kata/57a0e5c372292dd76d000d7e
func repeatString(n int, s string) string {
    var result string
    for i := 0; i < n; i++ {
        result += s
    }
    return result
}
