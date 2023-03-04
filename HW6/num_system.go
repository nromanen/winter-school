package main

import (
	"fmt"
	"strconv"
)

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func main() {
	var a int
	var b int

	fmt.Println("please, enter a (num)")
	fmt.Scan(&a)
	fmt.Println("please, enter b (number system)")
	fmt.Scan(&b)

	v, _ := ConvertInt("10", a, b)
	fmt.Println("convert(a, b) - >", v)

	// 12(10) -> C(16)
	// 11(10) -> B(16)
}
