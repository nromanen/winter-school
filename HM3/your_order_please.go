package kata
//https://www.codewars.com/kata/55c45be3b2079eccff00010f
import (
    "sort"
    "strconv"
    "strings"
)

func order(words string) string {
    if words == "" {
        return ""
    }
    wordList := strings.Split(words, " ")
    sortedList := make([]string, len(wordList))
    for _, word := range wordList {
        num, _ := strconv.Atoi(strings.TrimFunc(word, func(r rune) bool {
            return !unicode.IsNumber(r)
        }))
        sortedList[num-1] = word
    }
    return strings.Join(sortedList, " ")
}
