package utils

import (
	"strings"
	"unicode/utf8"
)

func Reverse(str string) string {
	var result strings.Builder
	result.Grow(len(str))
	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		result.WriteRune(r)
		str = str[:len(str)-size]
	}
	return result.String()
}
