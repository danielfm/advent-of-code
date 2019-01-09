package main

import (
	"fmt"
	"testing"
)

func Test_makeCircle(t *testing.T) {
	c := makeCircle(0)

	if c.next != c {
		t.Errorf("Expected next marble to be itself, but was %+v", c.next)
	}

	if c.prev != c {
		t.Errorf("Expected previous marble to be itself, but was %+v", c.prev)
	}

	if c.id != 0 {
		t.Errorf("Expected marble ID to be zero, but was %d", c.id)
	}
}

func Test_insert(t *testing.T) {
	c := makeCircle(0)

	m1 := c.insert(1)
	m2 := m1.insert(2)

	if c.id != 0 {
		t.Errorf("Expected marble ID to be 0, but was %d", c.id)
	}

	if c.prev.id != 2 {
		t.Errorf("Expected previous marble from 0 to be 2, but was %d", c.prev.id)
	}

	if c.next.id != 1 {
		t.Errorf("Expected previous marble from 0 to be 1, but was %d", c.next.id)
	}

	if m1.id != 1 {
		t.Errorf("Expected marble ID to be 1, but was %d", m2.id)
	}

	if m1.prev.id != 0 {
		t.Errorf("Expected previous marble from 1 to be 0, but was %d", m1.prev.id)
	}

	if m1.next.id != 2 {
		t.Errorf("Expected previous marble from 1 to be 2, but was %d", m1.next.id)
	}

	if m2.id != 2 {
		t.Errorf("Expected marble ID to be 2, but was %d", m2.id)
	}

	if m2.prev.id != 1 {
		t.Errorf("Expected previous marble from 2 to be 1, but was %d", m2.prev.id)
	}

	if m2.next.id != 0 {
		t.Errorf("Expected next marble from 2 to be 0, but was %d", m2.next.id)
	}
}

func Test_shift(t *testing.T) {
	c := makeCircle(0)
	c.insert(1).insert(2)

	if c.id != 0 {
		t.Errorf("Expected marble ID to be 0, but was %d", c.id)
	}

	// Clockwise
	c = c.shift(1)
	if c.id != 1 {
		t.Errorf("Expected marble ID to be 1, but was %d", c.id)
	}

	c = c.shift(1)
	if c.id != 2 {
		t.Errorf("Expected marble ID to be 2, but was %d", c.id)
	}

	c = c.shift(1)
	if c.id != 0 {
		t.Errorf("Expected marble ID to be 0, but was %d", c.id)
	}

	// Counter-clockwise
	c = c.shift(-1)
	if c.id != 2 {
		t.Errorf("Expected marble ID to be 2, but was %d", c.id)
	}

	c = c.shift(-1)
	if c.id != 1 {
		t.Errorf("Expected marble ID to be 1, but was %d", c.id)
	}

	c = c.shift(-1)
	if c.id != 0 {
		t.Errorf("Expected marble ID to be 0, but was %d", c.id)
	}
}

func Test_remove(t *testing.T) {
	c := makeCircle(0).insert(1).insert(2)
	c = c.remove()

	if c.id != 0 {
		t.Errorf("Expected marble ID to be 0, but was %d", c.id)
	}

	if c.next.id != 1 {
		t.Errorf("Expected marble ID next to the current one to be 1, but was %d", c.next.id)
	}

	if c.next.next.id != 0 {
		t.Errorf("Expected marble ID next to the next one to be 0, but was %d", c.next.id)
	}
}

func Test_playGame(t *testing.T) {
	testCases := []struct {
		numPlayers, lastMarble, expectedScore int
	}{
		{
			numPlayers:    9,
			lastMarble:    25,
			expectedScore: 32,
		},
		{
			numPlayers:    10,
			lastMarble:    1618,
			expectedScore: 8317,
		},
		{
			numPlayers:    13,
			lastMarble:    7999,
			expectedScore: 146373,
		},
		{
			numPlayers:    17,
			lastMarble:    1104,
			expectedScore: 2764,
		},
		{
			numPlayers:    21,
			lastMarble:    6111,
			expectedScore: 54718,
		},
		{
			numPlayers:    30,
			lastMarble:    5807,
			expectedScore: 37305,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d numPlayers %d lastMarble", tc.numPlayers, tc.lastMarble), func(t *testing.T) {
			actual := playGame(tc.numPlayers, tc.lastMarble)
			if actual != tc.expectedScore {
				t.Errorf("Expected winning score to be %d, but was %d", tc.expectedScore, actual)
			}
		})

	}
}
