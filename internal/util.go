package internal

import (
	"unicode"
)

func CountLeadingSpaces(s string) int {
	count := 0
	for _, ch := range s {
		if !unicode.IsSpace(ch) {
			break
		}
		count++
	}
	return count
}
