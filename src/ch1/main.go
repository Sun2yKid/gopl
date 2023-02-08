package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("GOPL: 你好，世界！")

	a := int(123)
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))

	var i int
	var u uint
	var up uintptr

	fmt.Printf("i Type:%T Size:%d\n", i, unsafe.Sizeof(i))
	fmt.Printf("u Type:%T Size:%d\n", u, unsafe.Sizeof(u))
	fmt.Printf("up Type:%T Size:%d\n", up, unsafe.Sizeof(up))
}
