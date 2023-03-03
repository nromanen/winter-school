package main

import (
	"fmt"
	"strconv"
)

func digitize(n uint) []uint {
	s := strconv.Itoa(int(n))
	var digits []uint
	for _, v := range s {
		d, _ := strconv.Atoi(string(v))
		digits = append(digits, uint(d))
	}
	return digits
}

func ComputeDepth(n uint) uint {
	digits := map[uint]bool{}
	var i uint
	for i = 1; len(digits) < 10; i++ {
		for _, d := range digitize(n * i) {
			digits[d] = true
		}
	}
	return i - 1
}

func main() {
	// expected - 9
	fmt.Println(ComputeDepth(42))
	// expected - 12
	fmt.Println(ComputeDepth(8))
	// expected - 36
	fmt.Println(ComputeDepth(25))
}
