package ciphers_test

import (
	"go-tasks/ciphers"
	"testing"
)

func TestCaesarEncode(t *testing.T) {
	str := "abcdef"

	res := ciphers.Caesar(str, 1)

	if *res != "xyzabc" {
		t.Error("Caesar should encode a simple example correctly")
	}
}

func TestCaesarDecode(t *testing.T) {
	str := "xyzabc"

	res := ciphers.Caesar(str, 0)

	if *res != "abcdef" {
		t.Error("Caesar should decode a simple example correctly")
	}
}

func TestCaesarEncodeUppercase(t *testing.T) {
	str := "ABCDEF"

	res := ciphers.Caesar(str, 1)

	if *res != "XYZABC" {
		t.Error("Caesar should encode uppercase letters correctly")
	}
}

func TestCaesarDecodeLowercase(t *testing.T) {
	str := "xyzabc"

	res := ciphers.Caesar(str, 0)

	if *res != "abcdef" {
		t.Error("Caesar should deocde uppercase letters correctly")
	}
}
