package kata

import "strings"

func Order(s string) string {
	split := strings.Split(s, " ")
	result := make([]string, len(split))
	for _, i := range split {
		for _, j := range i {
			if j >= '0' && j <= '9' {
				result[int(j-'1')] = i
				break
			}
		}
	}
	return strings.Join(result, " ")
}
