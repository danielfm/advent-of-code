package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var input = flag.String("input", "input", "Puzzle input file")

var claimExpr = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

type claim struct {
	id, x, y, width, height int
}

func main() {
	flag.Parse()

	claims := loadClaims(*input)

	fmt.Printf("Number of square inches with 2+ claims: %d\n", sumConflictingSquareInches(claims))
	fmt.Printf("Unconflicting claim: %d\n", findUnconflictingClaimID(claims))
}

func loadClaims(filename string) []claim {
	claims := []claim{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		claims = append(claims, parseClaim(scanner.Text()))
	}

	return claims
}

func parseClaim(claimStr string) claim {
	match := claimExpr.FindStringSubmatch(claimStr)

	id, err := strconv.Atoi(match[1])
	if err != nil {
		log.Fatal(err)
	}

	x, err := strconv.Atoi(match[2])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(match[3])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(match[4])
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(match[5])
	if err != nil {
		log.Fatal(err)
	}

	return claim{
		id:     id,
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func makeFabric(size int) [][]int {
	fabric := make([][]int, size)
	for i := range fabric {
		fabric[i] = make([]int, size)
	}
	return fabric
}

func sumConflictingSquareInches(claims []claim) int {
	fabric := makeFabric(1000)
	conflicting := 0

	for _, claim := range claims {
		for i := claim.x; i < claim.x+claim.width; i++ {
			for j := claim.y; j < claim.y+claim.height; j++ {
				if fabric[j][i] == 1 {
					conflicting += 1
				}
				fabric[j][i] += 1
			}
		}
	}

	return conflicting
}

func findUnconflictingClaimID(claims []claim) int {
	fabric := makeFabric(1000)

	for _, claim := range claims {
		for i := claim.x; i < claim.x+claim.width; i++ {
			for j := claim.y; j < claim.y+claim.height; j++ {
				if fabric[j][i] == 0 {
					fabric[j][i] = claim.id
				} else if fabric[j][i] > 0 {
					fabric[j][i] = -1
				}
			}
		}
	}

	for _, claim := range claims {
		untouched := true

		for i := claim.x; i < claim.x+claim.width; i++ {
			for j := claim.y; j < claim.y+claim.height; j++ {
				if fabric[j][i] != claim.id {
					untouched = false
				}
			}
		}

		if untouched {
			return claim.id
		}
	}

	return 0
}
