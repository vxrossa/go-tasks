package ciphers

import "errors"

type Direction string

type CaesarCipher struct{}

type CaesarActions interface {
	Handle(s string, encode int) (string, error)
	HandleLowercase(l byte, isEncode bool) string
	HandleUppercase(l byte, isEncode bool) string
}

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

func (c *CaesarCipher) HandleUppercase(l byte, isEncode bool) string {
	return string(moveChar(l, MAX_CHAR_CODE_UPPER, MIN_CHAR_CODE_UPPER, isEncode))
}

func (c *CaesarCipher) HandleLowercase(l byte, isEncode bool) string {
	return string(moveChar(l, MAX_CHAR_CODE_LOWER, MIN_CHAR_CODE_LOWER, isEncode))
}

func (c *CaesarCipher) Handle(s string, e int) (string, error) {
	if invalidEncode(e) {
		return "", errors.New("invalid encode")
	}

	res := ""

	for i := range s {
		if s[i] >= MIN_CHAR_CODE_UPPER && s[i] <= MAX_CHAR_CODE_UPPER {
			res += c.HandleUppercase(s[i], e == 1)
		} else if s[i] >= MIN_CHAR_CODE_LOWER && s[i] <= MAX_CHAR_CODE_LOWER {
			res += c.HandleLowercase(s[i], e == 1)
		} else {
			res += string(s[i])
		}

	}

	return res, nil
}
