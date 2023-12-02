package iteration

// given a character, repeat it n times; if n is omitted default to 5
func Repeat(character string, n ...int) string {
	var repeated string
	var times int

	// check if the number of repetitions is defined
	// otherwise default to 5
	if len(n) == 1 {
		times = n[0]
	} else {
		times = 5
	}

	for i := 0; i < times; i++ {
		repeated += character
	}

	return repeated
}
