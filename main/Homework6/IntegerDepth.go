package kata

//Task 4 https://www.codewars.com/kata/59b401e24f98a813f9000026
func ComputeDepth(n uint) uint {
	var d uint = 1
	m := map[uint]bool{}
	for ; len(m) < 10; d++ {
		for i := d * n; i != 0; i /= 10 {
			m[i%10] = true
		}
	}
	return d - 1
}