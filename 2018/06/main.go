package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	id, x, y int
}

type grid map[point]int

var input = flag.String("input", "input", "Puzzle input file")

func main() {
	flag.Parse()

	grid := loadGrid(*input)

	fmt.Printf("Size of the largest non-infinite area: %d.\n", grid.largestNonInfiniteArea())
	fmt.Printf("Size of the region containing all locations within total max distance: %d.\n", grid.areaWithin(10000))
}

func (g grid) maxPoint() point {
	maxX, maxY := 0, 0

	for point := range g {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	return point{x: maxX, y: maxY}
}

func loadGrid(filename string) grid {
	id := 0
	grid := make(grid)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pointData := strings.Split(scanner.Text(), ", ")

		x, err := strconv.Atoi(pointData[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(pointData[1])
		if err != nil {
			log.Fatal(err)
		}

		grid[point{id: id, x: x, y: y}], id = 0, id+1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func manhattanDistance(p1, p2 point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func (g grid) largestNonInfiniteArea() int {
	distGrid := make(grid)
	maxPoint := g.maxPoint()

	// Calculate the closest element for every position. The value for each
	// entry will be set to the ID of its closest point.
	for x := 0; x <= maxPoint.x; x++ {
		for y := 0; y <= maxPoint.y; y++ {
			gp := point{x: x, y: y}

			closestDistance := math.MaxInt64
			closestPoints := []point{}

			for p := range g {
				d := manhattanDistance(gp, p)

				if d < closestDistance {
					closestDistance = d
					closestPoints = []point{p}
				} else if d == closestDistance {
					closestPoints = append(closestPoints, p)
				}
			}

			// Disconsider points equally closest to two or more points.
			if len(closestPoints) > 1 {
				continue
			}

			distGrid[gp] = closestPoints[0].id
		}
	}

	// Compute the area for each point in the distance grid; and assume the
	// points closest to the edges have infinite area.
	area := map[int]int{}
	for gp, pid := range distGrid {
		if gp.y == 0 || gp.y == maxPoint.y || gp.x == 0 || gp.x == maxPoint.x {
			area[pid] = math.MaxInt64
		} else if area[pid] != math.MaxInt64 {
			area[pid]++
		}
	}

	pointIDWithLargestArea := -1
	for pid, a := range area {
		if a != math.MaxInt64 && (pointIDWithLargestArea == -1 || a > area[pointIDWithLargestArea]) {
			pointIDWithLargestArea = pid
		}
	}

	return area[pointIDWithLargestArea]
}

func (g grid) areaWithin(threshold int) int {
	sum := 0

	grid := make(grid)
	maxPoint := g.maxPoint()

	// Calculate the closest element for every position.
	for x := 0; x <= maxPoint.x; x++ {
		for y := 0; y <= maxPoint.y; y++ {
			p := point{x: x, y: y}
			sum := 0

			for pg := range g {
				sum += manhattanDistance(p, pg)
			}

			grid[p] = sum
		}
	}

	for x := 0; x <= maxPoint.x; x++ {
		for y := 0; y <= maxPoint.y; y++ {
			p := point{x: x, y: y}

			if grid[p] < threshold {
				sum++
			}
		}
	}

	return sum
}
