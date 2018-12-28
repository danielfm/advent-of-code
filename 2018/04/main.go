package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var input = flag.String("input", "input", "Puzzle input file")

var eventExpr = regexp.MustCompile(`^\[(\d{4}\-\d{2}\-\d{2} \d{2}:\d{2})\] (.+)$`)
var shiftExpr = regexp.MustCompile(`^Guard #(\d+) begins shift$`)

var tsFormat = "2006-01-02 15:04"

type eventType int

const (
	eventTypeAsleep eventType = iota
	eventTypeAwake
)

type event struct {
	time      time.Time
	eventType eventType
	guardID   int
}

func main() {
	flag.Parse()

	events := loadEvents(*input)
	sleepMatrix := computeSleepMatrix(events)

	guardID1 := findMostAsleepGuardID(sleepMatrix)

	fmt.Printf("Guard with most minutes asleep: #%d.\n", guardID1)
	fmt.Printf("Most asleep minute was %d.\n", findMostAsleepMinute(sleepMatrix[guardID1]))

	guardID2 := findGuardWithMostAsleepMinute(sleepMatrix)

	fmt.Printf("Guard with the most asleep minute: #%d.\n", guardID2)
	fmt.Printf("Most asleep minute was %d.\n", findMostAsleepMinute(sleepMatrix[guardID2]))
}

func loadEvents(filename string) []event {
	events := []event{}

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fileContent), "\n")
	sort.Strings(lines)

	currentGuardID := 0
	for _, line := range lines {
		match := eventExpr.FindStringSubmatch(line)
		if match == nil {
			log.Fatal("Unexpected non-match: ", line)
		}

		// New shift; keep track of the current Guard ID
		shiftMatch := shiftExpr.FindStringSubmatch(match[2])
		if shiftMatch != nil {
			if currentGuardID, err = strconv.Atoi(shiftMatch[1]); err != nil {
				log.Fatal(err)
			}
		} else if match[2] == "falls asleep" {
			events = append(events, event{
				time:      parseTimestamp(match[1]),
				eventType: eventTypeAsleep,
				guardID:   currentGuardID,
			})

		} else if match[2] == "wakes up" {
			events = append(events, event{
				time:      parseTimestamp(match[1]),
				eventType: eventTypeAwake,
				guardID:   currentGuardID,
			})
		} else {
			log.Fatal("Unknown event: ", match[2])
		}
	}

	return events
}

func computeSleepMatrix(events []event) map[int][]int {
	sleepMatrix := map[int][]int{}

	for i := 0; i <= len(events)-2; i += 2 {
		guardID := events[i].guardID

		asleepEvent := events[i]
		awakeEvent := events[i+1]

		if _, ok := sleepMatrix[guardID]; !ok {
			sleepMatrix[guardID] = make([]int, 61)
		}

		// The remaining elements correspond to the minutes asleep in that particular minute (1-59)
		for j := asleepEvent.time.Minute(); j < awakeEvent.time.Minute(); j++ {
			sleepMatrix[guardID][j] += 1
		}

		// Last element contains the total number of minutes asleep
		sleepMatrix[guardID][60] += (awakeEvent.time.Minute() - asleepEvent.time.Minute())
	}

	return sleepMatrix
}

func findMostAsleepGuardID(sleepMatrix map[int][]int) int {
	guardID := 0

	for i, guardSleepTimes := range sleepMatrix {
		if guardID == 0 || guardSleepTimes[60] > sleepMatrix[i][60] {
			guardID = i
		}
	}

	return guardID
}

func findGuardWithMostAsleepMinute(sleepMatrix map[int][]int) int {
	guardID, mostAsleepMinute := 0, 0

	for i, guardSleepTimes := range sleepMatrix {
		guardMostAsleepMinute := 0

		if guardID == 0 {
			guardID = i
		}

		// Find the most asleep minute for the current guard ID
		for j := 0; j < 60; j++ {
			if guardSleepTimes[j] > guardSleepTimes[guardMostAsleepMinute] {
				guardMostAsleepMinute = j
			}
		}

		// Update guard ID if the current guard has a minute with most times asleep
		if guardSleepTimes[guardMostAsleepMinute] > sleepMatrix[guardID][mostAsleepMinute] {
			guardID, mostAsleepMinute = i, guardMostAsleepMinute
		}
	}

	return guardID
}

func findMostAsleepMinute(guardSleepInfo []int) int {
	mostSleepedMinute := 0

	for i := 0; i < 60; i++ {
		if guardSleepInfo[i] >= guardSleepInfo[mostSleepedMinute] {
			mostSleepedMinute = i
		}
	}

	return mostSleepedMinute
}

func parseTimestamp(tsStr string) time.Time {
	ts, err := time.Parse(tsFormat, tsStr)
	if err != nil {
		log.Fatal(err)
	}
	return ts
}
