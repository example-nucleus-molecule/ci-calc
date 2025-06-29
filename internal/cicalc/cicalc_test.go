package cicalc

import (
	"math/big"
	"testing"
)

func TestPeriodic(t *testing.T) {
	testCases := []struct {
		name      string
		principal float64
		rate      float64
		term      float64
		freq      Frequency
		expected  int
	}{
		{
			name:      "with monthly frequency",
			principal: 10000,
			rate:      5,
			term:      5,
			freq:      Monthly,
			expected:  12834,
		}, {
			name:      "with quarterly frequency",
			principal: 10000,
			rate:      5,
			term:      5,
			freq:      Quarterly,
			expected:  12820,
		}, {
			name:      "with annual frequency",
			principal: 10000,
			rate:      5,
			term:      5,
			freq:      Annual,
			expected:  12763,
		}, {
			name:      "with atMaturity frequency",
			principal: 10000,
			rate:      5,
			term:      5,
			freq:      AtMaturity,
			expected:  12500,
		}, {
			name:      "with a fractional term",
			principal: 29200,
			rate:      7.79,
			term:      3 + getMonth(5), // 3 years + 5 months
			freq:      Quarterly,
			expected:  38007,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Periodic(tc.principal, tc.rate, tc.term, tc.freq)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func getMonth(month int) float64 {
	res, _ := big.NewRat(int64(month), 12).Float64()
	return res
}
