package main

import (
	"fmt"
	"regexp"
)

func checkJson(f string) bool {
	regex := `[{\[]{1}([,:{}\[\]0-9.\-+Eaeflnr-u \n\r\t]|".*?")+[}\]]{1}`
	check, err := regexp.MatchString(regex, f)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return check
}

func main() {

	fmt.Println(checkJson("{\n      \"id\": 1,\n      \"title\": \"Олещук Інна Володимирівна\",\n      \"label\": \"Oleschuk Inna\",\n      \"description\": \"08.08.1938-13.05.2001\",\n      \"image\": \"photos/f.png\",\n      \"itemTitleColor\": \"#ffc0cb\"\n    },\n    {\n      \"id\": 2,\n      \"title\": \"Олещук  Віктор Андрійович\",\n      \"label\": \"Oleschuk Viktor\",\n      \"description\": \"27.02.1933 - 19.04.2005\",\n      \"image\": \"photos/b.png\",\n      \"itemTitleColor\": \"#c8e4fb\"\n    }"))

	fmt.Println(checkJson("{[\"object\": \"string\""))
}
