package txt

import (
	"strconv"
	"strings"
)

// Numeric removes non-numeric characters from a string and returns the result.
func Numeric(s string) string {
	if s == "" {
		return ""
	}

	sep := '.'

	if c := strings.Count(s, "."); c == 0 || c > 1 {
		sep = ','
	}

	// Remove invalid characters.
	s = strings.Map(func(r rune) rune {
		if r == sep {
			return '.'
		} else if r == '-' {
			return '-'
		} else if r < '0' || r > '9' {
			return -1
		}

		return r
	}, s)

	return s
}

// Float64 converts a string to a 64-bit floating point number or 0 if invalid.
func Float64(s string) float64 {
	if s == "" {
		return 0
	}

	f, err := strconv.ParseFloat(Numeric(s), 64)

	if err != nil {
		return 0
	}

	return f
}

// Int64 converts a string to a signed 64-bit integer or 0 if invalid.
func Int64(s string) int64 {
	if s == "" {
		return 0
	}

	i := strings.SplitN(Numeric(s), ".", 2)

	result, err := strconv.ParseInt(i[0], 10, 64)

	if err != nil {
		return 0
	}

	return result
}
