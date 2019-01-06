package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var input = flag.String("input", "input", "Puzzle input file")

type node struct {
	header   []int
	children []node
	meta     []int
}

func main() {
	flag.Parse()

	n := loadTree(*input)

	fmt.Printf("Sum of all metadata entries: %d\n", n.sumMetadata())
	fmt.Printf("Value of the root node: %d\n", n.value())
}

func loadTree(filename string) node {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	strData := strings.Split(string(fileContent), " ")
	data := make([]int, len(strData))

	for i, elem := range strData {
		n, ok := strconv.Atoi(elem)
		if ok != nil {
			log.Fatal(err)
		}
		data[i] = n
	}

	node, _ := parseNode(data, 0)
	return node
}

func parseNode(data []int, offset int) (node, int) {
	n := node{}

	n.header = data[offset : offset+2]
	offset += 2

	if n.header[0] > 0 {
		for i := 0; i < n.header[0]; i++ {
			child, newOffset := parseNode(data, offset)
			n.children = append(n.children, child)
			offset = newOffset
		}
	}

	n.meta = data[offset : offset+n.header[1]]
	offset += n.header[1]

	return n, offset
}

func (n node) sumMetadata() int {
	acc := 0

	for _, i := range n.meta {
		acc += i
	}

	for _, c := range n.children {
		acc += c.sumMetadata()
	}

	return acc
}

func (n node) value() int {
	value := 0

	if len(n.children) == 0 {
		return n.sumMetadata()
	}

	for _, i := range n.meta {
		if i < 1 || i > len(n.children) {
			continue
		}
		value += n.children[i-1].value()
	}

	return value
}
