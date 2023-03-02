package main

import (
	"fmt"
	"strconv"
)

func main() {

	v, _ := ConvertInt("111", 2, 16)
	fmt.Printf(" base 2|%s : base 16|%s\n", "111", v)

	v, _ = ConvertInt("1C", 16, 10)
	fmt.Printf("base 16|%s : base 10|%s\n", "1C", v)

	v, _ = ConvertInt("12", 8, 16)
	fmt.Printf("base 8|%s : base 16|%s\n", "12", v)

	v, _ = ConvertInt("11", 10, 8)
	fmt.Printf("base 10|%s : base 8|%s\n", "11", v)

}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
