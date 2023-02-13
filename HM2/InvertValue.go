package kata

func Invert(s []int) []int {
	for i := range s {
		s[i] = -s[i]
	}
	return s
}
