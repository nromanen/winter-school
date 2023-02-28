package kata

import (
	"strconv"
	"strings"
)

func Order(sentence string) string {

	var numberFind []string
	var stringSplit []string
	stringSplit = strings.Split(sentence, " ")
	stringOrder := make([]string, len(stringSplit))

	for i := 0; i < len(stringSplit); i++ {

		numberFind = strings.Split(stringSplit[i], "")

		for m := 0; m < len(numberFind); m++ {

			if _, err := strconv.Atoi(numberFind[m]); err == nil {
				num, _ := strconv.Atoi(numberFind[m])
				stringOrder[num-1] = stringSplit[i]
			}
		}
		sentence = strings.Join(stringOrder, " ")
	}

	return sentence
}
