package main

import (
	"fmt"
	"sort"
	"strings"
)

func inArray(a1 []string, a2 []string) []string {
	var result []string
	for _, s1 := range a1 {
		for _, s2 := range a2 {
			if strings.Contains(s2, s1) {
				result = append(result, s1)
				break
			}
		}
	}
	sort.Strings(result)
	return result
}

func main() {
	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	result := inArray(a1, a2)
	fmt.Println(result)
}
