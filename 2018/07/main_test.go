package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_loadGraph(t *testing.T) {
	graph := loadGraph("input.test")

	if !reflect.DeepEqual(graph.g["C"], []string{"A", "F"}) {
		t.Errorf("Expected node C to be connected with A and F, but it is connectd to %+v", graph.g["C"])
	}

	if !reflect.DeepEqual(graph.g["A"], []string{"B", "D"}) {
		t.Errorf("Expected node A to be connected with B and D, but it is connected to %+v", graph.g["A"])
	}

	if !reflect.DeepEqual(graph.g["B"], []string{"E"}) {
		t.Errorf("Expected node B to be connected with E, but it is connected to %+v", graph.g["B"])
	}

	if !reflect.DeepEqual(graph.g["D"], []string{"E"}) {
		t.Errorf("Expected node D to be connected with E, but it is connected to %+v", graph.g["D"])
	}

	if !reflect.DeepEqual(graph.g["F"], []string{"E"}) {
		t.Errorf("Expected node F to be connected with E, but it is connected to %+v", graph.g["F"])
	}

	if !reflect.DeepEqual(graph.g["E"], []string{}) {
		t.Errorf("Expected node E to not be connected with any other node, but it is connected to %+v", graph.g["E"])
	}
}

func Test_findPredecessors(t *testing.T) {
	testCases := []struct {
		node     string
		expected []string
	}{
		{
			node:     "C",
			expected: []string{},
		},
		{
			node:     "A",
			expected: []string{"C"},
		},
		{
			node:     "F",
			expected: []string{"C"},
		},
		{
			node:     "B",
			expected: []string{"A"},
		},
		{
			node:     "D",
			expected: []string{"A"},
		},
		{
			node:     "E",
			expected: []string{"B", "D", "F"},
		},
	}

	graph := loadGraph("input.test")

	for _, tc := range testCases {
		t.Run(tc.node, func(t *testing.T) {
			actual := graph.findPredecessors(tc.node)

			sort.Sort(sort.StringSlice(actual))
			sort.Sort(sort.StringSlice(tc.expected))

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected predecessors of %s to be %+v, but was %+v", tc.node, tc.expected, actual)
			}
		})
	}
}

func Test_stepDuration(t *testing.T) {
	testCases := []struct {
		stepName string
		duration int
		expected int
	}{
		{
			stepName: "A",
			duration: 60,
			expected: 61,
		},
		{
			stepName: "B",
			duration: 60,
			expected: 62,
		},
		{
			stepName: "Z",
			duration: 60,
			expected: 86,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.stepName, func(t *testing.T) {
			actual := stepDuration(tc.duration, tc.stepName)
			if actual != tc.expected {
				t.Errorf("Expected duration for %s is %d, but was %d", tc.stepName, tc.expected, actual)
			}
		})
	}
}

func Test_execute(t *testing.T) {
	graph := loadGraph("input.test")

	expectedSteps := []string{"C", "A", "B", "F", "D", "E"}
	expectedTime := 15

	steps, time := graph.execute(0, 2)

	if !reflect.DeepEqual(steps, expectedSteps) {
		t.Errorf("Expected completion time to be %d, but was %d", expectedSteps, steps)
	}

	if time != expectedTime {
		t.Errorf("Expected completion time to be %d, but was %d", expectedTime, time)
	}
}
