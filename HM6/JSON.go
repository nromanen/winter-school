package main

//За допомогою регулярного виразу перевірити, чи заданий рядок є представленням json-формату (без масивів)

import (
	"fmt"
	"regexp"
)

func isJSON(s string) bool {
	re := regexp.MustCompile(`^\s*\{(?:\s*"[^"\n]*"\s*:\s*(?:"[^"\n]*"|null|true|false|\d+(?:\.\d+)?)(?:,\s*"[^"\n]*"\s*:\s*(?:"[^"\n]*"|null|true|false|\d+(?:\.\d+)?))*\s*)?\}\s*$`)
	return re.MatchString(s)
}

func main() {
	fmt.Println(isJSON("{\n  \"firstName\": \"Ok\",\n  \"lastName\": \"Trni\",\n  \"isAlive\": true,\n  \"age\": 15,\n  \"address\": {\n    \"streetAddress\": \"21 2nd Street\",\n    \"city\": \"New York\",\n    \"state\": \"NY\",\n    \"postalCode\": \"10021-3100\"\n  },\n  \"phoneNumbers\": [\n    {\n      \"type\": \"home\",\n      \"number\": \"212 555-1234\"\n    },\n    {\n      \"type\": \"office\",\n      \"number\": \"646 555-4567\"\n    }\n  ],\n  \"children\": [\n      \"Catherine\",\n      \"Thomas\",\n      \"Trevor\"\n  ],\n  \"spouse\": null\n}"))
	fmt.Println(isJSON("{[\"object\": \"string\""))
}
