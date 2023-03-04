package kata

func RepeatStr(repetitions int, value string) string {
	var repeat string

	for i := 0; i < repetitions; i++ {
		repeat += value
	}

	return repeat
}
