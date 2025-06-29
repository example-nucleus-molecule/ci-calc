// Package cicalc implements a compound interest calculator.
package cicalc

import (
	"math"
)

type Frequency int

const (
	Monthly    Frequency = 12
	Quarterly  Frequency = 4
	Annual     Frequency = 1
	AtMaturity Frequency = 0
)

// Periodic takes a principal amount, an interest rate (in percentage points,
// e.g. 1.25 for 1.25%), a term (in years), and an interest frequency, and
// returns a final amount (rounded to the nearest whole number).
func Periodic(
	principal float64, rate float64, term float64, freq Frequency,
) int {
	floatFreq := float64(freq)
	return int(math.Round(principal * math.Pow(1+((rate/100)/floatFreq), term*floatFreq)))
}
