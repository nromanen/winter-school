package main

//Написати функцію Add(n), яка буде повертати функцію, що додає до свого аргументу n

//f := Add(5)

//f(10) -> 15

import "fmt"

func Add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {
	f := Add(5)
	result := f(10)
	fmt.Println(result) // 15
}
