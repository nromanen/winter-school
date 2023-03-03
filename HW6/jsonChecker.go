package main

import (
	"fmt"
	"regexp"
)

func checkJSON(f string) bool {
	regex := `[{\[]{1}([,:{}\[\]0-9.\-+Eaeflnr-u \n\r\t]|".*?")+[}\]]{1}`
	match, err := regexp.MatchString(regex, f)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return match
}

func main() {
	// expected - true (case: https://en.wikipedia.org/wiki/JSON)
	fmt.Println(checkJSON("{\n  \"firstName\": \"John\",\n  \"lastName\": \"Smith\",\n  \"isAlive\": true,\n  \"age\": 27,\n  \"address\": {\n    \"streetAddress\": \"21 2nd Street\",\n    \"city\": \"New York\",\n    \"state\": \"NY\",\n    \"postalCode\": \"10021-3100\"\n  },\n  \"phoneNumbers\": [\n    {\n      \"type\": \"home\",\n      \"number\": \"212 555-1234\"\n    },\n    {\n      \"type\": \"office\",\n      \"number\": \"646 555-4567\"\n    }\n  ],\n  \"children\": [\n      \"Catherine\",\n      \"Thomas\",\n      \"Trevor\"\n  ],\n  \"spouse\": null\n}"))
	// expected - false
	fmt.Println(checkJSON("{[\"object\": \"string\""))
}
