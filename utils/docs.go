package utils

import (
	"unicode"
)

func toInt(r rune) int {
	return int(r - '0')
}

func allDigit(doc string) bool {
	for _, r := range doc {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}
