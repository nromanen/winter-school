package main

import (
	"fmt"
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	set := map[string]bool{}
	var result []string
	for _, sub := range array1 {
		for _, word := range array2 {
			if strings.Contains(word, sub) {
				set[sub] = true
				break
			}
		}
	}
	for k := range set {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}

func main() {
	fmt.Println(InArray([]string{"arp", "live", "strong"},
		[]string{"lively", "alive", "harp", "sharp", "armstrong"}))
	fmt.Println(InArray([]string{"tarp", "mice", "bull"},
		[]string{"lively", "alive", "harp", "sharp", "armstrong"}))
}
