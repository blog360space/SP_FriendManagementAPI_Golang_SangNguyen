package helpers

import (
	"math/rand"
	"regexp"
)

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

func AddItemToArray(arraySrc []string, item string) []string {
	Block{
		Try: func() {
			arraySrc = append(arraySrc, item)
		},
		Catch: func(e Exception) {

		},
	}.Do()
	return arraySrc
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

