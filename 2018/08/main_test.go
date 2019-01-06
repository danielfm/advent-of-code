package main

import (
	"reflect"
	"testing"
)

func Test_loadTree(t *testing.T) {
	var n node
	var expectedMeta []int

	tree := loadTree("input.test")

	// Root node (A)
	n = tree

	if len(n.children) != 2 {
		t.Errorf("Expected node A to have 2 child nodes, but it had %d", len(n.children))
	}

	expectedMeta = []int{1, 1, 2}
	if !reflect.DeepEqual(n.meta, expectedMeta) {
		t.Errorf("Expected node A to have meta %+v, but was %+v", expectedMeta, n.meta)
	}

	// Node B
	n = tree.children[0]

	if len(n.children) != 0 {
		t.Errorf("Expected node B to have 0 child nodes, but it had %d", len(n.children))
	}

	expectedMeta = []int{10, 11, 12}
	if !reflect.DeepEqual(n.meta, expectedMeta) {
		t.Errorf("Expected node B to have meta %+v, but was %+v", expectedMeta, n.meta)
	}

	// Node C
	n = tree.children[1]

	if len(n.children) != 1 {
		t.Errorf("Expected node C to have 1 child nodes, but it had %d", len(n.children))
	}

	expectedMeta = []int{2}
	if !reflect.DeepEqual(n.meta, expectedMeta) {
		t.Errorf("Expected node C to have meta %+v, but was %+v", expectedMeta, n.meta)
	}

	// Node D
	n = tree.children[1].children[0]

	if len(n.children) != 0 {
		t.Errorf("Expected node D to have 0 child nodes, but it had %d", len(n.children))
	}

	expectedMeta = []int{99}
	if !reflect.DeepEqual(n.meta, expectedMeta) {
		t.Errorf("Expected node D to have meta %+v, but was %+v", expectedMeta, n.meta)
	}
}

func Test_sumMetadata(t *testing.T) {
	tree := loadTree("input.test")
	s := tree.sumMetadata()

	if s != 138 {
		t.Errorf("Sum of metadata expected to be %d, but was %d", 138, s)
	}
}

func Test_value(t *testing.T) {
	tree := loadTree("input.test")
	s := tree.value()

	if s != 66 {
		t.Errorf("Value of root node expected to be %d, but was %d", 66, s)
	}
}
