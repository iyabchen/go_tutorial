package main

import (
	"fmt"
)

func main() {

	a := 1
	var p *int = &a

	fmt.Println(*p)

	if a := 10; a > 1 {
		// b is only valid inside the if block
		// if same name variable from outside, then the outside one will be ignored
		fmt.Println(a)
	}

	for {
		a++
		if a > 3 {
			break // break for
		}
		fmt.Println(a)
	}
	fmt.Println("over")

	for a <= 5 {
		a++
		fmt.Println(a)
	}

	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a)
	}

	vara := 0
	switch vara {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a>0")
		fallthrough
	default:
		fmt.Println("none")
	}

LABEL1:
	for {
		for i := 0; i < 3; i++ {
			if i > 1 {
				break LABEL1 // break to the same level of LABEL1. c.f goto
			}
		}
	}
	fmt.Println("break suc")

LABEL2:

	for i := 0; i < 3; i++ {
		for {
			fmt.Println(i)
			continue LABEL2
		}
	}
	fmt.Println("fin")

	// LABEL3:
	// 	for i := 0; i < 3; i++ {
	// 		for {
	// 			fmt.Println(i)
	// 			goto LABEL3 // cause deadlock
	// 		}
	// 	}
}
