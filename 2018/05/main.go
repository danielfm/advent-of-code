package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

var input = flag.String("input", "input", "Puzzle input file")

func main() {
	flag.Parse()

	polymer := loadPolymer(*input)

	output := processPolymer(polymer)
	fmt.Printf("The resulting polymer has %d units.\n", len(output))

	shortestPolymer := findShortestPolymer(polymer)
	fmt.Printf("Shortest polymer has %d units.\n", len(shortestPolymer))
}

func loadPolymer(filename string) string {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(fileContent)
}

func stripUnits(polymer string, unit rune) string {
	output := []rune{}
	unit1, unit2 := unicode.ToUpper(unit), unicode.ToLower(unit)

	for _, u := range polymer {
		if u != unit1 && u != unit2 {
			output = append(output, u)
		}
	}

	return string(output)
}

func processPolymer(polymer string) string {
	output := []rune{}

	for _, u := range polymer {
		if len(output) > 0 && doesUnitsReact(u, output[len(output)-1]) {
			output = output[0 : len(output)-1]
		} else {
			output = append(output, u)
		}
	}

	return string(output)
}

func findShortestPolymer(polymer string) string {
	shortestPolymer := polymer

	for u := 'a'; u <= 'z'; u++ {
		strippedPolymer := stripUnits(polymer, u)
		result := processPolymer(strippedPolymer)

		if len(result) < len(shortestPolymer) {
			shortestPolymer = result
		}
	}

	return shortestPolymer
}

func doesUnitsReact(u1, u2 rune) bool {
	return u1 != u2 && unicode.ToUpper(u1) == unicode.ToUpper(u2)
}
