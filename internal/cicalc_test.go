package cicalc

import "testing"

func TestPeriodic(t *testing.T) {
	testCases := []struct {
		name      string
		principal int
		rate      float32
		term      int
		freq      Frequency
		expected  int
	}{
		{
			name:      "success",
			principal: 10000,
			rate:      5,
			term:      5,
			freq:      Quarterly,
			expected:  12820,
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
