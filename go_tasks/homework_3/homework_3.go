package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func reverseString(s string) {
	var result string = ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(rune(s[i]))
	}
	fmt.Println(result)
}

func createPhoneNumber(numbers [10]uint) {
	var result string = "("
	for i := 0; i < 3; i++ {
		result += strconv.FormatUint(uint64(numbers[i]), 10)
	}
	result += ") "
	for i := 3; i < 6; i++ {
		result += strconv.FormatUint(uint64(numbers[i]), 10)
	}
	result += "-"
	for i := 6; i < 10; i++ {
		result += strconv.FormatUint(uint64(numbers[i]), 10)
	}
	fmt.Println(result)
}

func capitalizeString(s string) {
	substrings := strings.SplitN(s, "_", -1)
	if len(substrings) == 1 {
		substrings = strings.SplitN(s, "-", -1)
	}

	var result string = ""
	for ind := range substrings {
		if ind == 0 {
			result += substrings[ind]
		} else {
			result += strings.Title(substrings[ind])
		}
	}
	fmt.Println(result)
}

func RepeatStr(repetitions int, value string) {
	var result string = ""
	for i := 0; i < repetitions; i++ {
		result += value
	}
	fmt.Println(result)
}

func orderString(sentence string) {
	//if(sentence=="" || sentence==" "){return sentence}
	substrings := strings.SplitN(sentence, " ", -1)
	var result string = ""
	re := regexp.MustCompile("[0-9]+")

	for i := 1; i <= len(substrings); i++ {
		for ind, substr := range substrings {
			var num = strconv.Itoa(i)
			if num == re.FindAllString(substrings[ind], -1)[0] {
				result += substr
				if i < len(substrings) {
					result += " "
				}
				break
			}
		}
	}
	fmt.Println(result)
}

func main() {
	reverseString("hello world!")
	createPhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	capitalizeString("the-stealth-warrior")
	RepeatStr(2, "hello")
	orderString("is2 Thi1s T4est 3a")
}
