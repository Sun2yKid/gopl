// 6.2.1 Nil is a valid recevier value
package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	fmt.Println(m) // map[lang:[en]]

	m.Add("item", "1")
	fmt.Println(m) // map[item:[1] lang:[en]]

	m.Add("item", "2")
	fmt.Println(m) // map[item:[1 2] lang:[en]]

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"        first value
	fmt.Println(m["item"])     // "[1, 2]"	  direct map access

	m = nil

	fmt.Println(m)         //map[]
	fmt.Println(m["item"]) // []
	m.Add("item", "3")     // panic: assignment to entry in nil map
}
