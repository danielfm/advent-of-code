package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_loadGrid(t *testing.T) {
	points := []point{
		{id: 0, x: 1, y: 1},
		{id: 1, x: 1, y: 6},
		{id: 2, x: 8, y: 3},
		{id: 3, x: 3, y: 4},
		{id: 4, x: 5, y: 5},
		{id: 5, x: 8, y: 9},
	}

	expected := make(grid)
	for _, p := range points {
		expected[p] = 0
	}

	actual := loadGrid("input.test")

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected grid to be %+v, but was %+v", expected, actual)
	}
}

func Test_maxPoint(t *testing.T) {
	grid := loadGrid("input.test")

	expected := point{x: 8, y: 9}
	actual := grid.maxPoint()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected max point to be %+v, but was %+v", expected, actual)
	}
}

func Test_distanceTo(t *testing.T) {
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
			actual := tc.p1.distanceTo(tc.p2)
			if actual != tc.expected {
				t.Errorf("Expected distance to be %d, but was %d", tc.expected, actual)
			}
		})
	}
}

func Test_largestNonInfiniteArea(t *testing.T) {
	grid := loadGrid("input.test")

	expected := 17
	actual := grid.largestNonInfiniteArea()

	if actual != expected {
		t.Errorf("Expected largest non-infinite area to be %d, but was %d", expected, actual)
	}
}

func Test_areaWithin(t *testing.T) {
	grid := loadGrid("input.test")

	expected := 16
	actual := grid.areaWithin(32)

	if actual != expected {
		t.Errorf("Expected area within max total distance to be %d, but was %d", expected, actual)
	}
}
