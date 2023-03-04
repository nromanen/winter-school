package main

import (
	"fmt"
	"sort"
	"strings"
)

func InArray(a1 []string, a2 []string) []string {
	result := []string{}

	uniqStr := make(map[string]struct{})
	for _, s := range a1 {
		uniqStr[s] = struct{}{}
	}

	var a1Uniq = []string{}
	for key, _ := range uniqStr {
		a1Uniq = append(a1Uniq, key)
	}

	for _, str1 := range a1Uniq {
		for _, str2 := range a2 {
			if strings.Count(str2, str1) > 0 {
				result = append(result, str1)
				break
			}
		}
	}

	sort.Strings(result)

	return result

}

func main() {

	fmt.Println("Task 3 :", InArray([]string{"God", "Jesus", "Devil"}, []string{"God", "is", "Devil"}))
}
