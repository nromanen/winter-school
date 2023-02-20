package kata

// https://www.codewars.com/kata/550554fd08b86f84fe000a58/train/go

import (
	"sort"
	"strings"
)

func CheckIn(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func InArray(array1 []string, array2 []string) []string {

	result := []string{}
	for _, str1 := range array1 {
		for _, str2 := range array2 {
			if strings.Contains(str2, str1) && !CheckIn(str1, result) {
				result = append(result, str1)
			}
		}
	}
	sort.Strings(result)
	return result
}
