package lsproduct

import (
	"errors"
	"regexp"
	"strconv"
)

func charToDigit(c byte) (digit int) {
	digit, _ = strconv.Atoi(string(c))
	return
}

func productOf(sequence string, start, span int) (prod int) {
	prod = charToDigit(sequence[start])
	for i := start + 1; i < start+span; i++ {
		prod *= charToDigit(sequence[i])
	}
	return
}

func validateParams(sequence string, span int) (bool, error) {
	size := len(sequence)
	if span < 0 {
		return false, errors.New("Span must be positive")
	}
	if span > size {
		return false, errors.New("Span can't exceed the length of the digits")
	}
	if size == 0 || span == 0 {
		return false, nil
	}
	if re := regexp.MustCompile("\\D"); re.MatchString(sequence) {
		return false, errors.New("Invalid character in the sequence")
	}
	return true, nil
}

// LargestSeriesProduct calculates the largest product of a series
// of length span in the given sequence of digits
func LargestSeriesProduct(sequence string, span int) (int, error) {
	size := len(sequence)
	res, err := validateParams(sequence, span)
	if err != nil {
		return -1, err
	}
	if !res {
		return 1, nil
	}

	largest := 0
	for i := 0; i <= size-span; i++ {
		if cur := productOf(sequence, i, span); cur > largest {
			largest = cur
		}
	}
	return largest, nil
}
