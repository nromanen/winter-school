package main
// task 3
import "fmt"

func Add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}
