package bob

import (
	"regexp"
	s "strings"
)

const meaningless string = " \n\r\t"

func isYelling(remark string) bool {
	hasLetters := regexp.MustCompile(`[A-Za-z]+`).MatchString
	return hasLetters(remark) && s.ToUpper(remark) == remark
}

func isQuestioning(remark string) bool {
	return s.HasSuffix(s.Trim(remark, meaningless), "?")
}

func isSilent(remark string) bool {
	return len(s.Trim(remark, meaningless)) == 0
}

// Hey alternates Bob's answer given a remark
func Hey(remark string) string {
	if isYelling(remark) {
		return "Whoa, chill out!"
	} else if isQuestioning(remark) {
		return "Sure."
	} else if isSilent(remark) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
