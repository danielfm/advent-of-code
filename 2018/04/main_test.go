package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseTimestamp(t *testing.T) {
	expected := time.Date(1518, 11, 1, 0, 5, 0, 0, time.UTC)
	actual := parseTimestamp("1518-11-01 00:05")

	diff := expected.Sub(actual)
	if diff != 0 {
		t.Errorf("Expected diff between dates to be zero, but was %d", diff)
	}
}

func Test_loadEvents(t *testing.T) {
	expectedEvents := []event{
		{
			time:      parseTimestamp("1518-11-01 00:05"),
			guardID:   10,
			eventType: eventTypeAsleep,
		},
		{
			time:      parseTimestamp("1518-11-01 00:25"),
			guardID:   10,
			eventType: eventTypeAwake,
		},
		{
			time:      parseTimestamp("1518-11-01 00:30"),
			guardID:   10,
			eventType: eventTypeAsleep,
		},
		{
			time:      parseTimestamp("1518-11-01 00:55"),
			guardID:   10,
			eventType: eventTypeAwake,
		},
		{
			time:      parseTimestamp("1518-11-02 00:40"),
			guardID:   99,
			eventType: eventTypeAsleep,
		},
		{
			time:      parseTimestamp("1518-11-02 00:50"),
			guardID:   99,
			eventType: eventTypeAwake,
		},
		{
			time:      parseTimestamp("1518-11-03 00:24"),
			guardID:   10,
			eventType: eventTypeAsleep,
		},
		{
			time:      parseTimestamp("1518-11-03 00:29"),
			guardID:   10,
			eventType: eventTypeAwake,
		},
		{
			time:      parseTimestamp("1518-11-04 00:36"),
			guardID:   99,
			eventType: eventTypeAsleep,
		},
		{
			time:      parseTimestamp("1518-11-04 00:46"),
			guardID:   99,
			eventType: eventTypeAwake,
		},
		{
			time:      parseTimestamp("1518-11-05 00:45"),
			guardID:   99,
			eventType: eventTypeAsleep,
		},
		{
			time:      parseTimestamp("1518-11-05 00:55"),
			guardID:   99,
			eventType: eventTypeAwake,
		},
	}

	actualEvents := loadEvents("input.test")

	if len(expectedEvents) != len(actualEvents) {
		t.Errorf("Unexpected number of events: %d, but was %d", len(expectedEvents), len(actualEvents))
	}

	for i := range expectedEvents {
		expected, actual := expectedEvents[i], actualEvents[i]
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Unexpected event: %+v, but was %+v", expected, actual)
		}
	}
}

func Test_computeSleepMatrix(t *testing.T) {
	expected := map[int][]int{
		10: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 50},
		99: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 30},
	}

	events := loadEvents("input.test")
	actual := computeSleepMatrix(events)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Incorrect sleep matrix calculation: %+v, but was %+v", expected, actual)
	}
}

func Test_findMostAsleepGuardID(t *testing.T) {
	events := loadEvents("input.test")
	sleepMatrix := computeSleepMatrix(events)

	expected := 10
	actual := findMostAsleepGuardID(sleepMatrix)

	if expected != actual {
		t.Errorf("Incorrect guard ID: %d, but was %d", expected, actual)
	}
}

func Test_findGuardWithMostAsleepMinute(t *testing.T) {
	events := loadEvents("input.test")
	sleepMatrix := computeSleepMatrix(events)

	expected := 99
	actual := findGuardWithMostAsleepMinute(sleepMatrix)

	if expected != actual {
		t.Errorf("Incorrect guard ID: %d, but was %d", expected, actual)
	}
}

func Test_findMostAsleepMinute(t *testing.T) {
	events := loadEvents("input.test")
	sleepMatrix := computeSleepMatrix(events)

	expected := 24
	actual := findMostAsleepMinute(sleepMatrix[10])

	if expected != actual {
		t.Errorf("Incorrect guard ID: %d, but was %d", expected, actual)
	}
}
