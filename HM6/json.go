package main
// task 2
import "regexp"

func IsJSON(s string) bool {
    pattern := `^\s*\{\s*"(\w+)"\s*:\s*"(\w+)"(\s*,\s*"(\w+)"\s*:\s*"(\w+)")*\s*\}\s*$`
    matched, err := regexp.MatchString(pattern, s)
    return matched && err == nil
}
