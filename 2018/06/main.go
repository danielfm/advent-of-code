package main

import (
	"math"
)

type point struct {
	x, y int
}

func manhattanDistance(p1, p2 point) int {
	return int(math.Abs(float64(p1.x)-float64(p2.x)) + math.Abs(float64(p1.y)-float64(p2.y)))
}
