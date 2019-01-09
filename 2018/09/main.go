package main

import (
	"flag"
	"fmt"
)

var numPlayers = flag.Int("num-players", 425, "Number of players")
var lastMarble = flag.Int("last-marble", 70848, "Last marble worth, in points")

type marble struct {
	id         int
	prev, next *marble
}

func main() {
	flag.Parse()

	score := playGame(*numPlayers, *lastMarble)
	fmt.Printf("Winning score for %d players and %d last marble score: %d\n", *numPlayers, *lastMarble, score)
}

func playGame(numPlayers, lastMarble int) int {
	c := makeCircle(0)
	scores := make([]int, numPlayers)

	for marble := 1; marble <= lastMarble; marble++ {
		p := marble % numPlayers

		if marble%23 > 0 {
			c = c.shift(1).insert(marble)
		} else {
			c = c.shift(-7)
			scores[p] += c.id + marble
			c = c.remove()
		}
	}

	maxScore := 0
	for _, s := range scores {
		if s > maxScore {
			maxScore = s
		}
	}

	return maxScore
}

func makeCircle(id int) *marble {
	m := marble{
		id: id,
	}

	m.prev = &m
	m.next = &m

	return &m
}

func (m *marble) shift(times int) *marble {
	c := m
	clockwise := true

	if times < 0 {
		clockwise = false
		times *= -1
	}

	for n := 0; n < times; n++ {
		if clockwise {
			c = c.next
		} else {
			c = c.prev
		}
	}

	return c
}

func (m *marble) insert(id int) *marble {
	curr := marble{
		id: id,
	}

	curr.prev, curr.next = m, m.next
	m.next.prev, m.next = &curr, &curr

	return &curr
}

func (m *marble) remove() *marble {
	n, p := m.next, m.prev
	n.prev, p.next = p, n

	return n
}
