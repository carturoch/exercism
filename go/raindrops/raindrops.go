package raindrops

import "strconv"

// Convert returns the raindrops for the string or the
// the given number as string if no factors where drops.
func Convert(n int) string {
	res := ""
	if n%3 == 0 {
		res += "Pling"
	}
	if n%5 == 0 {
		res += "Plang"
	}
	if n%7 == 0 {
		res += "Plong"
	}
	if res != "" {
		return res
	}
	return strconv.Itoa(n)
}
