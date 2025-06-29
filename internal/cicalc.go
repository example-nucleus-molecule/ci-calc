// Package cicalc implements a compound interest calculator.
package cicalc

type Frequency int

const (
	Monthly Frequency = iota
	Quarterly
	Annually
	AtMaturity
)

// Periodic takes a principal amount, an interest rate, a term (in years), and
// an interest frequency, and returns a final amount (rounded to the nearest
// whole number).
func Periodic(principal int, rate float32, term int, freq Frequency) int {
	return 0
}
