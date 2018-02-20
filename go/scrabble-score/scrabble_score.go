package scrabble

import (
	"strings"
)

var scores = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func scoreForLetter(letter rune) int {
	for k, v := range scores {
		if strings.Contains(k, strings.ToUpper(string(letter))) {
			return v
		}
	}
	return 0
}

// Score calculates the scrabble value of the given word
func Score(w string) int {
	score := 0
	for _, letter := range w {
		score += scoreForLetter(letter)
	}
	return score
}
