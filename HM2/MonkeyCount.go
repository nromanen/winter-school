package kata

func monkeyCount(n int) []int {
	monkey := []int{}
	for i := 1; i <= n; i++ {
		monkey = append(monkey, i)
	}
	return monkey

}
