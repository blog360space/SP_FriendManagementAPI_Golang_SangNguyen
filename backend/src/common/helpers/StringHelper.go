package helpers

import (
	"reflect"
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

func ItemExists(arraySrc []string, item string) bool {
	var output = false
	Block{
		Try: func() {
			arr := reflect.ValueOf(arraySrc)

			for i := 0; i < arr.Len(); i++ {
				if arr.Index(i).Interface() == item {
					output = true
				}
			}

		},
		Catch: func(e Exception) {

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

