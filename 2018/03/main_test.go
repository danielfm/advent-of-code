package main

import (
	"reflect"
	"testing"
)

func Test_parseClaim(t *testing.T) {
	testCases := []struct {
		claimStr string
		expected claim
	}{
		{
			claimStr: "#1 @ 1,3: 4x4",
			expected: claim{id: 1, x: 1, y: 3, width: 4, height: 4},
		},
		{
			claimStr: "#2 @ 3,1: 4x4",
			expected: claim{id: 2, x: 3, y: 1, width: 4, height: 4},
		},
		{
			claimStr: "#3 @ 5,5: 2x2",
			expected: claim{id: 3, x: 5, y: 5, width: 2, height: 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.claimStr, func(t *testing.T) {
			actual := parseClaim(tc.claimStr)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got the wrong object: %+v, want %+v", actual, tc.expected)
			}
		})
	}
}

func Test_sumConflictingSquareInches(t *testing.T) {
	claims := []claim{
		{id: 1, x: 1, y: 3, width: 4, height: 4},
		{id: 2, x: 3, y: 1, width: 4, height: 4},
		{id: 3, x: 5, y: 5, width: 2, height: 2},
	}

	actual := sumConflictingSquareInches(claims)
	if actual != 4 {
		t.Errorf("got the wrong number of square inches with 2+ claims: %d, want %d", actual, 4)
	}
}

func Test_findUnconflictingClaimID(t *testing.T) {
	claims := []claim{
		{id: 1, x: 1, y: 3, width: 4, height: 4},
		{id: 2, x: 3, y: 1, width: 4, height: 4},
		{id: 3, x: 5, y: 5, width: 2, height: 2},
	}

	actual := findUnconflictingClaimID(claims)
	if actual != 3 {
		t.Errorf("got the wrong claim ID: %d, want %d", actual, 3)
	}
}
