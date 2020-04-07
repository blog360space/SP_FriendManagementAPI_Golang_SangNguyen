package helpers

import "regexp"

var emailRegex = regexp.MustCompile("\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*")

func IsEmail(input string) bool {
	var output = true
	Block{
		Try: func() {
			output = emailRegex.MatchString(input)
		},
		Catch: func(e Exception) {
			output = false
		},
	}.Do()
	return output
}