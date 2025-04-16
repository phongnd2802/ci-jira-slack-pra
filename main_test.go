package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Prime: 2", 2, true},
		{"Prime: 3", 3, true},
		{"Not Prime: 4", 4, false},
		{"Prime: 5", 5, true},
		{"Not Prime: 9", 9, false},
		{"Prime: 11", 11, true},
		{"Not Prime: 1", 1, false},
		{"Not Prime: 0", 0, false},
		{"Negative number", -7, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsPrime(tc.input)
			require.Equal(t, tc.expected, result, "input: %d", tc.input)
		})
	}
}
