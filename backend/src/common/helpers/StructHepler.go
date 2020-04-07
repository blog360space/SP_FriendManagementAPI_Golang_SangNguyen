package helpers

func IsNull(h interface{}) bool{
	var output = false
	Block{
		Try: func() {
			if h == nil {
				output = true
			}
		},
		Catch: func(e Exception) {
			output = true
		},
	}.Do()
	return output
}