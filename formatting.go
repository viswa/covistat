package main

import (
	"fmt"
	"strings"
)

// localize inserts commas to num based on Indian locale
func localize(num int) string {
	digits := []rune(fmt.Sprint(num))
	var builder strings.Builder
	var written int // no. of characters written to builder
	sep := 3        // no. of places between comma placement

	// digit characters are written to builder in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		if written == sep {
			builder.WriteString(",")
			// reset no. of written characters after each comma placed
			written = 0
			sep = 2
		}
		builder.WriteRune(digits[i])
		written++
	}

	// builder string is to be further reversed
	reversed := builder.String()
	digits = []rune(reversed)
	n := len(digits)
	for i := 0; i < n/2; i++ {
		digits[i], digits[n-1-i] = digits[n-1-i], digits[i]
	}
	return string(digits)
}
