package main

import (
	"fmt"
	"testing"
)

func Test_processPolymer(t *testing.T) {
	testCases := []struct {
		polymer  string
		expected string
	}{
		{
			polymer:  "aA",
			expected: "",
		},
		{
			polymer:  "abBA",
			expected: "",
		},
		{
			polymer:  "abAB",
			expected: "abAB",
		},
		{
			polymer:  "aabAAB",
			expected: "aabAAB",
		},
		{
			polymer:  "dabAcCaCBAcCcaDA",
			expected: "dabCBAcaDA",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.polymer, func(t *testing.T) {
			actual := processPolymer(tc.polymer)

			if string(actual) != tc.expected {
				t.Errorf("got the wrong output: %s, want %s", actual, tc.expected)
			}
		})
	}
}

func Test_stripUnits(t *testing.T) {
	testCases := []struct {
		polymer  string
		without  rune
		expected string
	}{
		{
			polymer:  "dabAcCaCBAcCcaDA",
			without:  'a',
			expected: "dbcCCBcCcD",
		},
		{
			polymer:  "dabAcCaCBAcCcaDA",
			without:  'b',
			expected: "daAcCaCAcCcaDA",
		},
		{
			polymer:  "dabAcCaCBAcCcaDA",
			without:  'c',
			expected: "dabAaBAaDA",
		},
		{
			polymer:  "dabAcCaCBAcCcaDA",
			without:  'd',
			expected: "abAcCaCBAcCcaA",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Polymer %s without unit %c", tc.polymer, tc.without), func(t *testing.T) {
			actual := stripUnits(tc.polymer, tc.without)

			if actual != tc.expected {
				t.Errorf("got the wrong output: %s, want %s", actual, tc.expected)
			}
		})
	}
}

func Test_findShortestPolymer(t *testing.T) {
	testCases := []struct {
		polymer  string
		expected string
	}{
		{
			polymer:  "dabAcCaCBAcCcaDA",
			expected: "daDA",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.polymer, func(t *testing.T) {
			actual := findShortestPolymer(tc.polymer)

			if actual != tc.expected {
				t.Errorf("got the wrong output: %s, want %s", actual, tc.expected)
			}
		})
	}
}
