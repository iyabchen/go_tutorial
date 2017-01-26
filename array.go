package main

import (
	"fmt"
)

func main() {
	var a [10]int
	// var b [1]int
	// b = a error

	c := [20]int{19: 1} // 19 is index, so 1 at the 19th index
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5} // implicitly define length
	fmt.Println(len(d))

	var p *[10]int = &a // a pointer pointing to a variable which is [100]int type
	fmt.Println(p)

	x, y := 1, 2
	var q = [...]*int{&x, &y} // an array variable which has 2 pointer, each pointer points to a int variable
	fmt.Println(*q[0])

	m := [2]int{1, 2}
	n := [2]int{1, 2}
	fmt.Println(m == n) // the comparison must be between the same type

	vara := [10]int{}
	vara[2] = 2

	pp := new([10]int)
	pp[2] = 2
	fmt.Println(*pp)

	multia := [...][3]int{ // only the first one can use ...
		{1, 1, 1},
		{2: 3}} // mind the ending } must be in the same line
	fmt.Println(multia)
}
