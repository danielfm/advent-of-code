package main

import (
	"fmt"
	"testing"
)

func Test_manhattanDistance(t *testing.T) {
	testCases := []struct {
		p1, p2   point
		expected int
	}{
		{
			p1:       point{x: 1, y: 1},
			p2:       point{x: 1, y: 6},
			expected: 5,
		},
		{
			p1:       point{x: 3, y: 4},
			p2:       point{x: 1, y: 6},
			expected: 4,
		},
		{
			p1:       point{x: 1, y: 6},
			p2:       point{x: 8, y: 9},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("p1=%+v,p2=%+v", tc.p1, tc.p2), func(t *testing.T) {
			actual := manhattanDistance(tc.p1, tc.p2)
			if actual != tc.expected {
				t.Errorf("Expected distance to be %d, but was %d", tc.expected, actual)
			}
		})
	}
}
