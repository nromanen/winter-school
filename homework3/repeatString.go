package main

import "fmt"

func repeatString(n int, s string) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += s
	}
	return repeated
}

func main() {

	fmt.Println(repeatString(6, "I"))     // prints "IIIIII"
	fmt.Println(repeatString(5, "Hello")) // prints "HelloHelloHelloHelloHello"

}
