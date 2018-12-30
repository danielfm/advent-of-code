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
	x, y int
}

var input = flag.String("input", "input", "Puzzle input file")

func main() {
	flag.Parse()

	points := loadPoints(*input)
	grid := makeGrid(points)
	area := calculateArea(grid)

	fmt.Printf("Size of the largest non-infinite area: %d.\n", pointIDWithLargestNonInfiniteArea(area))
}

func loadPoints(filename string) []point {
	points := []point{}

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

		points = append(points, point{
			x: x,
			y: y,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return points
}

func makeGrid(points []point) [][]int {
	maxX, maxY := 0, 0

	for _, point := range points {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	// Initialize the empty grid
	grid := make([][]int, maxY+1)
	for y := range grid {
		grid[y] = make([]int, maxX+1)
	}

	// Calculate the closest element for every position
	for y := range grid {
		for x := range grid[y] {
			closestDistance := math.MaxInt64
			closestPoints := []int{}

			for pi, p := range points {
				d := manhattanDistance(point{x: x, y: y}, p)

				if d < closestDistance {
					closestDistance = d
					closestPoints = []int{pi}
				} else if d == closestDistance {
					closestPoints = append(closestPoints, pi)
				}
			}

			if len(closestPoints) > 1 {
				grid[y][x] = -1
			} else {
				grid[y][x] = closestPoints[0]
			}
		}
	}

	return grid
}

func printGrid(grid [][]int) {
	for x := range grid {
		line := ""
		for y := range grid[x] {
			if grid[x][y] == math.MaxInt64 {
				line += ". "
			} else {
				line += fmt.Sprintf("%+d ", grid[x][y])
			}
		}
		fmt.Println(line)
	}
}

func calculateArea(grid [][]int) map[int]int {
	area := map[int]int{}

	for y := range grid {
		for x := range grid[y] {
			pid := grid[y][x]

			if pid == -1 {
				continue
			}

			if y == 0 || y == len(grid)-1 || x == 0 || x == len(grid[y])-1 {
				area[pid] = math.MaxInt64
			} else if area[pid] != math.MaxInt64 {
				area[pid]++
			}
		}
	}

	return area
}

func pointIDWithLargestNonInfiniteArea(area map[int]int) int {
	id := -1

	for pid := range area {
		if area[pid] != math.MaxInt64 && (id == -1 || area[pid] > area[id]) {
			id = pid
		}
	}

	return area[id]
}

func manhattanDistance(p1, p2 point) int {
	return int(math.Abs(float64(p1.x)-float64(p2.x)) + math.Abs(float64(p1.y)-float64(p2.y)))
}
