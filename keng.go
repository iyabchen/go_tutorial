package main

import (
	"fmt"
	"time"
)

func foo1(s []int) []int {
	s = append(s, 3) // if append, then it has to return the address, since it might be copied to a new address
	return s
}

func main() {
	s := make([]int, 0)
	fmt.Println(s)

	s1 := foo1(s) // s is passed using address. However, s was appended and copied to a new address
	fmt.Println(s)
	fmt.Println(s1)

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$

	t := time.Now()
	fmt.Println(t.Format(time.ANSIC)) // time must use const to format, never use a concrete time string

	str := []string{"a", "b", "c"}
	for _, v := range str {
		// go func() {
		// 	fmt.Println(v) // closure loop issue

		// }()
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(2 * time.Second)
}
