// https://www.codewars.com/kata/59a8570b570190d313000037/train/go

package kata

func SumCubes(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}
