package hamming

import (
	"errors"
)

// Distance calculates the Hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strands must have equal length")
	}
	distance := 0
	for i, nucleotide := range a {
		if byte(nucleotide) != b[i] {
			distance++
		}
	}

	return distance, nil
}
