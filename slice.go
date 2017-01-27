package main

import (
	"fmt"
)

func main() {
	var s1 []int // [...] or [number] means array, not slice
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 = a[5:len(a)] // including starting index, but not end index, meaning a[5 6 7 8 9]
	s2 := a[:5]
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := make([]int, 3, 10) //  like arraylist, if filled in more than 10, then the capcaity will double
	fmt.Println(s3, len(s3), cap(s3))

	sa := a[3:8]
	fmt.Println(sa, len(sa), cap(sa)) // capacity is the original array's len minus the index
	sb := sa[3:]                      // reslice
	fmt.Println(sb, len(sb), cap(sb))

	s4 := make([]int, 3, 6)
	fmt.Printf("%p\n", s4)

	s4 = append(s4, 3, 4, 5)
	fmt.Printf("%p %v\n", s4, s4) // addr not change

	s4 = append(s4, 3)
	fmt.Printf("%p %v\n", s4, s4) // addr changed, memory space reassigned

	s5 := a[2:5]
	s6 := a[1:3]
	fmt.Println(s5, s6)
	s5[0] = 100 // the referenced array is changed
	fmt.Println(s5, s6, cap(s6))

	s6 = append(s6, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	s5[0] = 200 // s6 has changed to another addr, so the change won't cover s6
	fmt.Println(s5, s6)

	s7 := []int{1, 2, 3, 4, 5, 6}
	s8 := []int{7, 8, 9}

	copy(s8, s7[4:6]) // use s7 to overwrite s8
	fmt.Println(s7, s8)

	shw := s7[:]
	fmt.Println(shw)

}
