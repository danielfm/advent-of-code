package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var input = flag.String("input", "input", "Puzzle input file")

var depExpr = regexp.MustCompile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin.$`)

type graph struct {
	g map[string][]string
}

func main() {
	var g graph
	flag.Parse()

	g = loadGraph(*input)
	steps, _ := g.execute(0, 1)
	fmt.Printf("Ordered steps: %s\n", strings.Join(steps, ""))

	g = loadGraph(*input)
	_, time := g.execute(60, 5)
	fmt.Printf("Completion time with 5 workers: %d\n", time)
}

func (g graph) addEdge(a, b string) {
	var ok bool

	_, ok = g.g[a]
	if !ok {
		g.g[a] = []string{}
	}

	_, ok = g.g[b]
	if !ok {
		g.g[b] = []string{}
	}

	g.g[a] = append(g.g[a], b)
}

func newGraph() graph {
	g := graph{}
	g.g = map[string][]string{}
	return g
}

func loadGraph(filename string) graph {
	g := newGraph()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match := depExpr.FindStringSubmatch(scanner.Text())
		g.addEdge(match[1], match[2])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return g
}

func (g graph) findPredecessors(node string) []string {
	pred := []string{}

	for prd, suc := range g.g {
		for _, s := range suc {
			if node == s {
				pred = append(pred, prd)
			}
		}
	}

	return pred
}

func (g graph) findRootNodes() []string {
	rootNodes := []string{}

	for n := range g.g {
		if len(g.findPredecessors(n)) == 0 {
			rootNodes = append(rootNodes, n)
		}
	}

	sort.Sort(sort.StringSlice(rootNodes))
	return rootNodes
}

func stepDuration(duration int, stepName string) int {
	return 1 + duration + int(stepName[0]-'A')
}

func (g graph) execute(duration, nWorkers int) ([]string, int) {
	seconds := 0
	runningTasks := map[string]int{}
	doneTasks := []string{}

	for {
		for t, d := range runningTasks {
			if stepDuration(duration, t) == seconds-d {
				doneTasks = append(doneTasks, t)
				delete(runningTasks, t)
				delete(g.g, t)
			}
		}

		availableTasks := g.findRootNodes()

		for _, task := range availableTasks {
			if len(runningTasks) == nWorkers {
				break
			}

			if _, ok := runningTasks[task]; !ok {
				runningTasks[task] = seconds
			}
		}

		if len(g.g) == 0 && len(runningTasks) == 0 {
			break
		}

		seconds++
	}

	return doneTasks, seconds
}
