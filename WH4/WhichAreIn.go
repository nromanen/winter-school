package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	r := inArray(a1, a2)
	fmt.Println(r)
}

func inArray(a1, a2 []string) []string {
	var res []string
	for _, s1 := range a1 {
		for _, s2 := range a2 {
			if strings.Contains(s2, s1) {
				res = append(res, s1)
				break
			}
		}
	}
	sort.Strings(res)
	return res
}
