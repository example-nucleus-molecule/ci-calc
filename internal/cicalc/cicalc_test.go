package cicalc

import "testing"

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
