package main

import "fmt"

func countMonkeys(n int) []int {
    monkeys := make([]int, n)
    for i := 1; i <= n; i++ {
        monkeys[i-1] = i
    }
    return monkeys
}

func main() {
    n := 5
    fmt.Println(countMonkeys(n))
}
