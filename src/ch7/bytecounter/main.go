package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0 // rest the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "Hello %s", name)
	fmt.Println(c)

}
