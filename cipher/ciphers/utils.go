package ciphers

func invalidEncode(encodeType int) bool {
	return encodeType != 1 && encodeType != 0
}

const (
	MIN_CHAR_CODE_LOWER = 97
	MAX_CHAR_CODE_LOWER = 122
	MIN_CHAR_CODE_UPPER = 65
	MAX_CHAR_CODE_UPPER = 90
)

const (
	LEFT  Direction = "left"
	RIGHT Direction = "right"
)
