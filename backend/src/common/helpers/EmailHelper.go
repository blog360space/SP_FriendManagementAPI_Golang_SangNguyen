package helpers

func ExtractEmails(input string) []string  {
	var output []string
	Block{
		Try: func() {
			matchAll := emailRegex.FindAllString(input, -1)
			for _, element := range matchAll {
				output = AddItemToArray(output, element)
			}
		},
		Catch: func(e Exception) {

		},
	}.Do()
	return output
}