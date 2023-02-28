package kata

//https://www.codewars.com/kata/59b401e24f98a813f9000026

func ComputeDepth(n uint) uint {
	var s uint = 1
	m := map[uint]bool{}
	for ; len(m) < 10; s++ {
		for i := s * n; i != 0; i /= 10 {
			m[i%10] = true
		}
	}
	return s - 1
}
