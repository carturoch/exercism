package lsproduct

import (
	"errors"
	"regexp"
	"strconv"
)

func charToDigit(c rune) (digit int) {
	digit, _ = strconv.Atoi(string(c))
	return
}

func productOf(sequence string, res chan<- int) {
	prod := 1
	for _, b := range sequence {
		prod *= charToDigit(b)
	}
	res <- prod
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

	upper := size - span
	products := make(chan int, upper)
	for i := 0; i <= upper; i++ {
		go productOf(sequence[i:i+span], products)
	}

	largest := 0
	for i := 0; i <= upper; i++ {
		if prod := <-products; prod > largest {
			largest = prod
		}
	}
	return largest, nil
}
