package main

import (
	"fmt"
)

func main() {
	a := [...]int{123, 45, 13, 67, 94, 23}
	fmt.Println(a)

	l := len(a)

	// insert sort
	// for i := 0; i < l; i++ {
	// 	for j := i; j > 0; j-- {
	// 		if a[j] < a[j-1] {
	// 			tmp := a[j-1]
	// 			a[j-1] = a[j]
	// 			a[j] = tmp
	// 		}
	// 	}
	// }

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if a[j] < a[i] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}

	fmt.Println(a)
}
