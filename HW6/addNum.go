<<<<<<< HEAD
package main

import "fmt"

func Add(x int) func(int) int {
	return func(i int) int {
		return x + i
	}
}

func main() {
	// expected - 15
	f := Add(5)
	fmt.Println(f(10))
	// expected - 10000
	g := Add(9999)
	fmt.Println(g(1))
}
=======
package main

import "fmt"

func Add(x int) func(int) int {
	return func(i int) int {
		return x + i
	}
}

func main() {
	// expected - 15
	f := Add(5)
	fmt.Println(f(10))
	// expected - 10000
	g := Add(9999)
	fmt.Println(g(1))
}
>>>>>>> 359a47e420f064ed1d3a826316ac80a49b2ae831
