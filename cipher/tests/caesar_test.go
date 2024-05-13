package ciphers_test

import (
	"go-tasks/ciphers"
	"testing"
)

func TestCaesarEncode(t *testing.T) {
	str := "abcdef"

	cipher := ciphers.CaesarCipher{}

	res, err := cipher.Handle(str, 1)

	if err != nil {
		t.Error(err)
	}

	if res != "xyzabc" {
		t.Error("Caesar should encode a simple example correctly")
	}
}

func TestCaesarDecode(t *testing.T) {
	str := "xyzabc"

	cipher := ciphers.CaesarCipher{}

	res, err := cipher.Handle(str, 0)

	if err != nil {
		t.Error(err)
	}

	if res != "abcdef" {
		t.Error("Caesar should decode a simple example correctly")
	}
}

func TestCaesarEncodeUppercase(t *testing.T) {
	str := "ABCDEF"

	cipher := ciphers.CaesarCipher{}

	res, err := cipher.Handle(str, 1)

	if err != nil {
		t.Error(err)
	}

	if res != "XYZABC" {
		t.Error("Caesar should encode uppercase letters correctly")
	}
}

func TestCaesarDecodeLowercase(t *testing.T) {
	str := "xyzabc"

	cipher := ciphers.CaesarCipher{}

	res, err := cipher.Handle(str, 0)

	if err != nil {
		t.Error(err)
	}

	if res != "abcdef" {
		t.Error("Caesar should deocde uppercase letters correctly")
	}
}
