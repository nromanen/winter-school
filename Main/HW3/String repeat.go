package kata

func RepeatStr(n int, s string) string {
var result string
for i := 0; i < n; i++ {
result += s
}
return result
}