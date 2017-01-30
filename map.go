package main

import (
	"fmt"
	"sort"
)

func main() {
	var m1 map[int]string
	m1 = make(map[int]string)

	m1[1] = "OK"
	fmt.Println(m1)
	a := m1[1]
	fmt.Println(a)
	delete(m1, 1)
	fmt.Println(m1)

	var m2 map[int]map[int]string = make(map[int]map[int]string) // only initiate the outside map, the inside one is not init
	m2[1] = make(map[int]string)
	m2[1][1] = "OK"
	fmt.Println(m2[1][1])

	value, ok := m2[2][1] // use the 2nd return value to judge whether it is initialized
	if !ok {
		m2[2] = make(map[int]string)
	}
	m2[2][1] = "GOOD"
	value, ok = m2[2][1]
	fmt.Println(value, ok)

	//iterate map object
	s1 := make([]map[int]string, 5) // a slice which has map as object
	for _, value := range s1 {      // the variable 'value' is just a value copy , not affecting content of slice
		value = make(map[int]string, 1)
		value[1] = "init_ok"
		fmt.Println(value)

	}
	fmt.Println(s1)

	for index, _ := range s1 {
		s1[index] = make(map[int]string, 1)
		s1[index][1] = "init_ok"
		fmt.Println(s1[index]) // the variable is just a value copy , not affecting content of slice

	}
	fmt.Println(s1)

	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

	for _, val := range s {
		fmt.Println(val, m[val])
	}

	m_rev := make(map[string]int)

	for k, v := range m {
		m_rev[v] = k
	}
	fmt.Println(m, m_rev)

}
