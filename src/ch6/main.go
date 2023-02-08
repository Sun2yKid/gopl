package main

import (
	"ch6/geometry"
	_ "ch6/intset"
	"fmt"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{4, 6}

	fmt.Println(geometry.Distance(p, q)) // function call
	fmt.Println(p.Distance(q))           // method call. expression `p.distance` is called a selector

	perim := geometry.Path{
		{X: 1, Y: 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())              // 12, method of geometry.Path
	fmt.Println(geometry.Path.Distance(perim)) // 12, standalone function

}
