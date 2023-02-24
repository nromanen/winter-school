package main

import (
	"fmt"
	"sort"
	"strings"
)

func inArray(a1 []string, a2 []string) []string {
	var r []string
	for _, s1 := range a1 {
		for _, s2 := range a2 {
			if strings.Contains(s2, s1) {
				r = append(r, s1)
				break
			}
		}
	}
	sort.Strings(r)
	return r
}

func main() {
	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	r := inArray(a1, a2)
	fmt.Println(r)

	a1 = []string{"tarp", "mice", "bull"}
	a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
	r = inArray(a1, a2)
	fmt.Println(r)
}
