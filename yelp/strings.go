package yelp

import "strconv"

// FloatString returns the string value of f.
func FloatString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// IntString returns the string value of i,
func IntString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// BoolString returns the bool value of b.
func BoolString(b bool) string {
	return strconv.FormatBool(b)
}
