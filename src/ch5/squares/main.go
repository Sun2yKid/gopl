package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
	g := squares()
	fmt.Println(g()) // "1"
	fmt.Println(g()) // "4"
	fmt.Println(g()) // "9"
	fmt.Println(g()) // "16"
}
