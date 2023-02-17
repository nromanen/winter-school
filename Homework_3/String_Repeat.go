package kata

func RepeatStr(repetitions int, value string) string {
	result := ""
	for i := 0; i < repetitions; i++ {
		result += value
	}

	return result
}
