package main

import (
	"ch6/geometry"
	_ "ch6/intset"
	_ "net/url"
	"fmt"
)

// An IntList is a linked list of integers
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail *IntList
}

// nil is a valid receiver value
// Sum returns the sum of list elements.
func (i *IntList) Sum() int {
	if i == nil {
		return 0
	} else {
		return i.Value + i.Tail.Sum()
	}
}


func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{4, 6}

	fmt.Println(geometry.Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))           // "5", method call. expression `p.distance` is called a selector

	perim := geometry.Path{
		{X: 1, Y: 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())              // 12, method of geometry.Path
	fmt.Println(geometry.Path.Distance(perim)) // 12, standalone function

	r := &geometry.Point{1, 2}
	fmt.Println(r) // &{1 2}
	r.ScaleBy(2)
	fmt.Println(*r) //{2 4}

}
