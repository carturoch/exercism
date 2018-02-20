package brackets

import (
	"strings"
)

const (
	openers = "[{("
	closers = "]})"
)

var pairs = map[byte]byte{']': '[', '}': '{', ')': '('}

// Bracket determines whether the brackets in str are balanced
func Bracket(str string) (bool, error) {
	acc := ""
	for _, c := range str {
		s := string(c)
		if strings.Contains(openers, s) {
			acc = s + acc
		}
		if strings.Contains(closers, s) {
			if acc == "" {
				return false, nil
			}
			if bracketsDoNotMatch(byte(c), acc[0]) {
				return false, nil
			}
			acc = acc[1:]
		}
	}
	return len(acc) == 0, nil
}

func bracketsDoNotMatch(a, b byte) bool {
	for close, open := range pairs {
		if a == close && b == open {
			return false
		}
	}
	return true
}
