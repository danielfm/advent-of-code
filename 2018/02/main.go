package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input = flag.String("input", "input", "Puzzle input file")

func main() {
	flag.Parse()

	boxIDs := loadBoxIDs(*input)

	log.Printf("Checksum is %d\n", calculateChecksum(boxIDs))
	fmt.Printf("Common letters: %s\n", commonLetters(boxIDs))
}

func loadBoxIDs(filename string) []string {
	boxIDs := []string{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boxIDs = append(boxIDs, scanner.Text())
	}

	return boxIDs
}

func calculateChecksum(boxIDs []string) int {
	accTwos, accThrees := 0, 0

	for _, boxID := range boxIDs {
		twos, threes := getOccurrences(boxID)
		accTwos, accThrees = accTwos+twos, accThrees+threes
	}

	return accTwos * accThrees
}

func getOccurrences(boxId string) (int, int) {
	twos, threes := 0, 0
	occurrences := map[rune]int{}

	for _, ch := range boxId {
		if count, ok := occurrences[ch]; ok {
			occurrences[ch] = count + 1
		} else {
			occurrences[ch] = 1
		}
	}

	for _, val := range occurrences {
		if twos == 0 && val == 2 {
			twos = twos + 1
		}
		if threes == 0 && val == 3 {
			threes = threes + 1
		}
	}

	return twos, threes
}

func commonLetters(boxIDs []string) string {
	for id1 := 0; id1 < len(boxIDs); id1++ {
		boxID1 := boxIDs[id1]

		for id2 := id1 + 1; id2 < len(boxIDs); id2++ {
			boxID2 := boxIDs[id2]
			commonLetters := []byte{}

			for cn := 0; cn < len(boxID1); cn++ {
				if boxID1[cn] == boxID2[cn] {
					commonLetters = append(commonLetters, boxID1[cn])
				}
			}

			// We found the correct boxes
			if len(commonLetters) == len(boxID1)-1 {
				return fmt.Sprintf("%s", commonLetters)
			}
		}
	}

	return ""
}
