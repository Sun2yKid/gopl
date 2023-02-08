package main

import (
	"ch6/geometry"
	"fmt"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{4, 6}

	fmt.Println(geometry.Distance(p, q)) // function call
	fmt.Println(p.Distance(q))           // method call

	perim := geometry.Path{
		{X: 1, Y: 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

}
