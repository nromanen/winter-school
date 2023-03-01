package main

import (
	"fmt"
	"sort"
	"strings"
)

func inArray(a1 []string, a2 []string) []string {
	// Create a set to store unique substrings
	set := make(map[string]bool)

	// Iterate through strings in a2 and check if they contain substrings from a1
	for _, s2 := range a2 {
		for _, s1 := range a1 {
			if strings.Contains(s2, s1) {
				set[s1] = true
			}
		}
	}

	// Create a slice of the substrings in lexicographical order
	r := make([]string, 0, len(set))
	for s := range set {
		r = append(r, s)
	}
	sort.Strings(r)

	return r
}

func main() {

	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(inArray(a1, a2)) // ["arp", "live", "strong"]

	a1 = []string{"tarp", "mice", "bull"}
	a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(inArray(a1, a2)) // []

}
