package kata
//https://www.codewars.com/kata/517abf86da9663f1d2000003
import "fmt"
func toCamelCase(s string) string {
    var res string
    var capitalizeNext bool
    for i := 0; i < len(s); i++ {
        if s[i] == '-' || s[i] == '_' {
            capitalizeNext = true
        } else {
            if capitalizeNext {
                res += strings.ToUpper(string(s[i]))
                capitalizeNext = false
            } else {
                res += string(s[i])
            }
        }
    }
    return res
}
