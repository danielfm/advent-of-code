package main

import (
	"testing"
)

func Test_calculateResultingFrequency(t *testing.T) {
	actual := calculateResultingFrequency([]int64{1, -2, 3, 1})

	if actual != 3 {
		t.Errorf("got the wrong resulting frequency: %d", actual)
	}
}

func Test_firstDuplicateFrequency(t *testing.T) {
	testCases := []struct {
		freqChanges []int64
		expected    int64
	}{
		{
			freqChanges: []int64{1, -1},
			expected:    0,
		},
		{
			freqChanges: []int64{3, 3, 4, -2, -4},
			expected:    10,
		},
		{
			freqChanges: []int64{-6, 3, 8, 5, -6},
			expected:    5,
		},
		{
			freqChanges: []int64{7, 7, -2, -7, -4},
			expected:    14,
		},
	}

	for _, tc := range testCases {
		actual := firstDuplicateFrequency(tc.freqChanges)
		if actual != tc.expected {
			t.Errorf("got the wrong first duplicate frequency: %d, want %d", actual, tc.expected)
		}
	}
}
