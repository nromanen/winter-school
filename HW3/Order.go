package kata

import (
	"regexp"
	"sort"
	"strings"
)

func Order(sentence string) string {
	if sentence == "" {
		return ""
	}
	w := strings.Split(sentence, " ")
	re := regexp.MustCompile("[1-9]")
	sort.Slice(w, func(w1, w2 int) bool {
		i := re.FindString(w[w1])
		j := re.FindString(w[w2])
		return i < j
	})
	return strings.Join(w, " ")
}
