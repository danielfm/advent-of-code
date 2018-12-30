package main

import (
	"fmt"
	"math"
	"reflect"
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

func Test_loadPoints(t *testing.T) {
	expected := []point{
		{
			x: 1,
			y: 1,
		},
		{
			x: 1,
			y: 6,
		},
		{
			x: 8,
			y: 3,
		},
		{
			x: 3,
			y: 4,
		},
		{
			x: 5,
			y: 5,
		},
		{
			x: 8,
			y: 9,
		},
	}
	actual := loadPoints("input.test")

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v to be %+v", actual, expected)
	}
}

func Test_makeGrid(t *testing.T) {
	points := loadPoints("input.test")

	expected := [][]int{
		{+0, +0, +0, +0, +0, -1, +2, +2, +2},
		{+0, +0, +0, +0, +0, -1, +2, +2, +2},
		{+0, +0, +0, +3, +3, +4, +2, +2, +2},
		{+0, +0, +3, +3, +3, +4, +2, +2, +2},
		{-1, -1, +3, +3, +3, +4, +4, +2, +2},
		{+1, +1, -1, +3, +4, +4, +4, +4, +2},
		{+1, +1, +1, -1, +4, +4, +4, +4, -1},
		{+1, +1, +1, -1, +4, +4, +4, +5, +5},
		{+1, +1, +1, -1, +4, +4, +5, +5, +5},
		{+1, +1, +1, -1, +5, +5, +5, +5, +5},
	}
	actual := makeGrid(points)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected grid to be %+v, but was %+v", expected, actual)
	}
}

func Test_calculateArea(t *testing.T) {
	points := loadPoints("input.test")
	grid := makeGrid(points)

	expected := map[int]int{
		0: math.MaxInt64,
		1: math.MaxInt64,
		2: math.MaxInt64,
		3: 9,
		4: 17,
		5: math.MaxInt64,
	}

	actual := calculateArea(grid)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected areas to be %+v, but was %+v", expected, actual)
	}
}

func Test_pointIDWithLargestNonInfiniteArea(t *testing.T) {
	points := loadPoints("input.test")
	grid := makeGrid(points)
	area := calculateArea(grid)

	expected := 17
	actual := pointIDWithLargestNonInfiniteArea(area)

	if actual != expected {
		t.Errorf("Expected largest non-infinite area to be %d, but was %d", expected, actual)
	}
}
