package kata

import (
	"strings"
)

func ToCamelCase(s string) string {
	s = strings.ReplaceAll(s, "-", ",")
	s = strings.ReplaceAll(s, "_", ",")
	s = strings.ReplaceAll(s, " ", ",")

	wd := strings.Split(s, ",")
	cml := wd[0]
	for i := 1; i < len(wd); i++ {
		cml += strings.Title(wd[i])
	}
	return cml
	// Lori-de-lori = Loridelori
}
