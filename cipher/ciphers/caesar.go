package ciphers

type Direction string

func moveChar(l byte, max int, min int, isEncode bool) byte {
	var letter int

	if isEncode {
		letter = int(l - 3)
	} else {
		letter = int(l + 3)
	}

	if letter < min {
		letter = max - (min - letter) + 1
	}
	if letter > max {
		letter = min + (letter - max) - 1
	}

	return byte(letter)
}

func uppercase(l byte, isEncode bool) string {
	return string(moveChar(l, MAX_CHAR_CODE_UPPER, MIN_CHAR_CODE_UPPER, isEncode))
}

func lowercase(l byte, isEncode bool) string {
	return string(moveChar(l, MAX_CHAR_CODE_LOWER, MIN_CHAR_CODE_LOWER, isEncode))
}

func Caesar(s string, e int) *string {
	encodeErr := invalidEncode(e)

	if encodeErr {
		panic("Wrong encode specified")
	}

	res := ""

	for i := range s {
		if s[i] >= MIN_CHAR_CODE_UPPER && s[i] <= MAX_CHAR_CODE_UPPER {
			res += uppercase(s[i], e == 1)
		} else if s[i] >= MIN_CHAR_CODE_LOWER && s[i] <= MAX_CHAR_CODE_LOWER {
			res += lowercase(s[i], e == 1)
		} else {
			res += string(s[i])
		}

	}

	return &res
}
