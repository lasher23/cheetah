package stringutil

import (
	"runtime"
	"strings"
)

func TrimReturn(string string) string {
	if runtime.GOOS == "windows" {
		return strings.TrimRight(string, "\r\n")
	} else {
		return strings.TrimRight(string, "\n")
	}
}

func ContainsIgnoreCase(a []string, x string) bool {
	for _, n := range a {
		if strings.EqualFold(x, n) {
			return true
		}
	}
	return false
}
