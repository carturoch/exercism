package luhn

import (
	"strconv"
	"strings"
)

// Valid determines whether a given number is valid
// according with Luhn algorithm
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	if len(input) < 2 {
		return false
	}
	sum := 0
	for i, char := range input {
		digit, _ := strconv.Atoi(string(char))
		if i%2 != 0 || i == 0 {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}
		sum += digit
	}
	return sum%10 == 0
}
