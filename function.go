package main

import (
	"fmt"
)

func main() {
	B("string", 1, 2)
	s1 := []int{1, 2, 3, 4}

	C(s1) // s1 is slice, so address copy is passed
	fmt.Println(s1)

	a, b := 1, 2
	D(a, b) // copied value is passed
	fmt.Println(a, b)

	c := 100
	E(&c)
	fmt.Println(c)

	tmp := F // function is also a type , tmp is a function now
	tmp()

	anonymousfunc := func() {
		fmt.Println("anonymous")
	}
	anonymousfunc()

	f := closure(10) // f is a function
	fmt.Printf("%p\n", f)
	fmt.Println(f(1))
	fmt.Println(f(2))

	fmt.Println("a)")
	defer fmt.Println("b")
	defer fmt.Println("c") // the later it define, the sooner it will be invoked

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) // i is referencing to the address, so not changing
		}() // meaning invoke this anonymous function
	}

	F()
	P()
	F()

}

func P() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")

		}
	}()
	panic("PANIC TEST IN P ")

}

func closure(x int) func(int) int {
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

func A() (int, string, int) {
	a, b, c := 1, "2", 3
	return a, b, c
}

func B(b string, a ...int) {
	fmt.Println(a)
}

func C(s []int) {
	s[0] = 5
	s[1] = 6
	fmt.Println(s)
}

func D(s ...int) {
	s[0] = 5
	s[1] = 6
	fmt.Println(s)
}

func E(a *int) {
	*a = 2
	fmt.Println(*a)
}

func F() {
	fmt.Println("Func F")
}
