package kata

import "fmt"

func SumCubes(n int) int {
sum := 0
for i := 1; i <= n; i++ {
sum += i * i * i
}
return sum
}

func main() {
fmt.Println(SumCubes(2)) // Output: 9
fmt.Println(SumCubes(3)) // Output: 36
}