package ciphers

import "errors"

type AtbashCipher struct{}

type AtbashActions interface {
	Handle(s string, encode int) (string, error)
	HandleUppercase(s string, isEncode bool) string
	HandleLowercase(s string, isEncode bool) string
}

func (a *AtbashCipher) HandleUppercase(l byte, isEncode bool) string {
	return string(MAX_CHAR_CODE_UPPER - (int(l) - MIN_CHAR_CODE_UPPER))
}

func (a *AtbashCipher) HandleLowercase(l byte, isEncode bool) string {
	return string(MAX_CHAR_CODE_LOWER - (int(l) - MIN_CHAR_CODE_LOWER))
}

func (a *AtbashCipher) Handle(s string, e int) (string, error) {
	if invalidEncode(e) {
		return "", errors.New("invalid encode")
	}

	res := ""

	for i := range s {
		if s[i] >= MIN_CHAR_CODE_UPPER && s[i] <= MAX_CHAR_CODE_UPPER {
			res += a.HandleUppercase(s[i], e == 1)
		} else if s[i] >= MIN_CHAR_CODE_LOWER && s[i] <= MAX_CHAR_CODE_LOWER {
			res += a.HandleLowercase(s[i], e == 1)
		} else {
			res += string(s[i])
		}
	}

	return res, nil
}
