package kata

func monkeyCount(n int) []int {
	monkey := []int{}
	for i := 0; i < n; i++ {
		monkey = append(monkey, i+1)
	}
	return monkey
}
