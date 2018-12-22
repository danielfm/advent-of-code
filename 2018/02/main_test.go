package main

import (
	"testing"
)

func Test_getOccurrences(t *testing.T) {
	testCases := []struct {
		boxId  string
		twos   int
		threes int
	}{
		{
			boxId:  "abcdef",
			twos:   0,
			threes: 0,
		},
		{
			boxId:  "bababc",
			twos:   1,
			threes: 1,
		},
		{
			boxId:  "abbcde",
			twos:   1,
			threes: 0,
		},
		{
			boxId:  "abcccd",
			twos:   0,
			threes: 1,
		},
		{
			boxId:  "aabcdd",
			twos:   1,
			threes: 0,
		},
		{
			boxId:  "abcdee",
			twos:   1,
			threes: 0,
		},
		{
			boxId:  "ababab",
			twos:   0,
			threes: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.boxId, func(t *testing.T) {
			twos, threes := getOccurrences(tc.boxId)
			if twos != tc.twos {
				t.Errorf("got the wrong number of twos: %d, want %d", twos, tc.twos)
			}
			if threes != tc.threes {
				t.Errorf("got the wrong number of threes: %d, want %d", threes, tc.threes)
			}
		})
	}
}

func Test_commonLetters(t *testing.T) {
	boxIDs := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	expected := "fgij"

	actual := commonLetters(boxIDs)
	if actual != expected {
		t.Errorf("got the wrong letters: %s, want %s", actual, expected)

	}
}
