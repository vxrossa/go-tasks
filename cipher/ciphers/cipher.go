package ciphers

import (
	"errors"
	"strconv"
	"strings"
)

type Cipher struct {
	Text    string
	Pattern string
	Result  string
	atbash  AtbashCipher
	caesar  CaesarCipher
	rot13   Rot13Cipher
}

func New(text string, pattern string) Cipher {
	return Cipher{
		Text:    text,
		Pattern: pattern,
	}
}

func (c *Cipher) Handle() (string, error) {
	opts := strings.Split(c.Pattern, "-")

	c.caesar = CaesarCipher{}
	c.atbash = AtbashCipher{}
	c.rot13 = Rot13Cipher{}

	for i := range opts {
		pattern := strings.Split(opts[i], "")

		encodeType, err := strconv.Atoi(pattern[1])

		if err != nil {
			panic("Could not define if the method is encoding or decoding")
		}

		switch pattern[0] {
		case "C":
			{
				textRes, err := c.caesar.Handle(c.Text, encodeType)

				if err != nil {
					return "", err
				}
				c.Text = textRes
			}
		case "A":
			{
				textRes, err := c.atbash.Handle(c.Text, encodeType)

				if err != nil {
					return "", err
				}
				c.Text = textRes
			}
		default:
			return "", errors.New("wrong cipher specified. Should be either C for Caesar, A for Atbash or R for ROT13")
		}
	}

	return c.Text, nil
}
