package strings

import (
	"strings"
	"unicode"
)

// IsUpperU 此方法的缺點，當遇到2字節的字母時可能會不如預期
func IsUpperU(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func IsLowerU(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	return strings.ToLower(s) == s
}
