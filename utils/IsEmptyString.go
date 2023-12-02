package utils

import "strings"

func IsEmptyString(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true
	}

	return false
}