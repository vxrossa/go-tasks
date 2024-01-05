package ciphers

import (
	"strconv"
	"strings"
)

type Cipher struct {
}

func (c Cipher) Handle(text *string, pattern string) *string {
	opts := strings.Split(pattern, "-")

	for i := range opts {
		pat := strings.Split(opts[i], "")

		encodeType, err := strconv.Atoi(pat[1])

		if err != nil {
			panic("Could not define if the method is encoding or decoding")
		}

		switch pat[0] {
		case "C":
			{
				text = Caesar(*text, encodeType)
			}
		case "A":
			{
				text = Atbash(text, encodeType)
			}
		case "R":
			{
				text = Rot13(text, encodeType)
			}
		default:
			panic("Wrong cipher specified. Should be either C for Caesar, A for Atbash or R for ROT13")
		}
	}

	return text
}
