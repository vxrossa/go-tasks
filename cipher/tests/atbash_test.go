package ciphers

import (
	"go-tasks/ciphers"
	"testing"
)

func TestAtbashEncodeUppercase(t *testing.T) {
	str := "ABCDEF"

	cipher := ciphers.AtbashCipher{}

	res, err := cipher.Handle(str, 1)

	if err != nil {
		t.Error(err)
	}

	if res != "ZYXWVU" {
		t.Error("Atbash should encode uppercase letters correctly")
	}
}

func TestAtbashEncodeLowercase(t *testing.T) {
	str := "abcdef"

	cipher := ciphers.AtbashCipher{}

	res, err := cipher.Handle(str, 1)

	if err != nil {
		t.Error(err)
	}

	if res != "zyxwvu" {
		t.Error("Atbash should encode lowercase letters correctly")
	}
}

func TestAtbashDecodeLowercase(t *testing.T) {
	str := "zyxwvu"

	cipher := ciphers.AtbashCipher{}

	res, err := cipher.Handle(str, 0)

	if err != nil {
		t.Error(err)
	}

	if res != "abcdef" {
		t.Error("Atbash should decode lowercase letters correctly")
	}
}

func TestAtbashDecodeUppercase(t *testing.T) {
	str := "ZYXWVU"

	cipher := ciphers.AtbashCipher{}

	res, err := cipher.Handle(str, 0)

	if err != nil {
		t.Error(err)
	}

	if res != "ABCDEF" {
		t.Error("Atbash should decode uppercase letters correctly")
	}
}
