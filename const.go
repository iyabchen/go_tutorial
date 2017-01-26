package main

import (
	"fmt"
)

const a, b, c = 1, "2", 'C'
const (
	d    = a + 2
	e    // not define, then use the above variable value
	f    = len(b)
	g, h = 1, "2"
	// i cannot just define one
	i, j
)

const (
	vara = 'A'
	varb
	varc = iota
	vard
)

const PI = 3.14

func main() {

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println(vara, varb, varc, vard) // A A 2 3
}
