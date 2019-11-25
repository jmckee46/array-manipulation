package main

import (
	"fmt"
	"testing"
)

// expected output is 200
func TestArrayManipulation1(t *testing.T) {
	queries := [][]int32{
		[]int32{1, 2, 100},
		[]int32{2, 5, 100},
		[]int32{3, 4, 100},
	}
	n := int32(5)
	total := arrayManipulation(n, queries)

	if total != 200 {
		t.Errorf("got %d instead of 200", total)
	}
	fmt.Println("total:", total)
}
