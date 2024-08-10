package alphanimeric

import (
	"regexp"
	"unicode"
)

func isAlphanumeric(input string) bool {
	if len(input) == 0 {
		return false
	}

	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

func alphanumeric(s string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return r.MatchString(s)
}

func IsAlphanumeric(str string) bool {
	return regexp.MustCompile("^[[:alnum:]]+$").MatchString(str)
}
