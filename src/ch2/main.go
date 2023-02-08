package main

import (
	"fmt"

	"ch2/popcount"
	"ch2/tempconv"
)

func main() {
	fmt.Println("ch2.main running...")

	c := tempconv.F2C(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))

	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.C2F(tempconv.BoilingC))

	popcountResult := popcount.PopCount(10)
	fmt.Println("%", popcountResult)
}
