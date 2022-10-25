package strTools

import "strings"

func RemoveRedundantWhiteSpaceChars(s ...*string) {
	for _, ss := range s {
		*ss = strings.Join(strings.Fields(*ss), " ")
	}
}

func RemoveSpecificChars(chars []string, s ...*string) {
	for _, ss := range s {
		RemoveRedundantWhiteSpaceChars(ss)
		for _, char := range chars {
			*ss = strings.Replace(*ss, char, "", -1)
		}
	}
}
