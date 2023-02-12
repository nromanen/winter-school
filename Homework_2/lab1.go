package kata

func Invert(arr []int) (invArr []int) {
	for _, value := range arr {
		invArr = append(invArr, -value)
	}
	return invArr
}
