package main

import (
	"fmt"
	"regexp"
)

func main() {

	json_string := `{
  "firstName": "John",
  "lastName": "Smith",
  "isAlive": true,
  "age": 27,
  "address": {
    "streetAddress": "21 2nd Street",
    "city": "New York",
    "state": "NY",
    "postalCode": "10021-3100"
  },
  "phoneNumbers": [
    {
      "type": "home",
      "number": "212 555-1234"
    },
    {
      "type": "office",
      "number": "646 555-4567"
    }
  ],
  "children": [
      "Catherine",
      "Thomas",
      "Trevor"
  ],
  "spouse": null
}`

	pattern := `[{\[]{1}([,:{}\[\]0-9.\-+Eaeflnr-u \n\r\t]|".*?")+[}\]]{1}`

	match, err := regexp.MatchString(pattern, json_string)

	if err != nil {
		fmt.Println("Error:", err)
	} else if match {
		fmt.Println("The string is a valid JSON format ")
	} else {
		fmt.Println("The string is not a valid JSON format ")
	}
}
