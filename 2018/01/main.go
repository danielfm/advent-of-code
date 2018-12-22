package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input = flag.String("input", "input", "Puzzle input file")

func main() {
	flag.Parse()

	freqChanges := readFrequencyChanges(*input)

	freq := calculateResultingFrequency(freqChanges)
	fmt.Printf("The resulting frequency is %d.\n", freq)

	dupFreq := firstDuplicateFrequency(freqChanges)
	fmt.Printf("First duplicate resulting frequency is %d.\n", dupFreq)
}

func readFrequencyChanges(filename string) []int64 {
	freqList := make([]int64, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		freqList = append(freqList, change)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return freqList
}

func calculateResultingFrequency(freqChanges []int64) int64 {
	var acc int64
	for _, freqChange := range freqChanges {
		acc = acc + freqChange
	}
	return acc
}

func firstDuplicateFrequency(freqChanges []int64) int64 {
	var acc int64
	occurrences := map[int64]bool{0: true}

	for {
		for _, freqChange := range freqChanges {
			acc = acc + freqChange

			if _, ok := occurrences[acc]; ok {
				return acc
			} else {
				occurrences[acc] = true
			}
		}
	}
}
