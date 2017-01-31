package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B int // even if B is type int

func main() {
	a := A{}
	a.Print() // even if defined with pointer, does not have to use *a.Print
	fmt.Println(a.Name)

	var b B
	b.Print()      // method value
	(*B).Print(&b) // &b is receiver, method expression

	var bb B
	bb = 0
	(*B).Increase(&bb)
	fmt.Println(bb)
}

func (a *A) Print() { // defined a method belonged to struct A, a *A is called a receiver
	a.Name = "AA" // pointer was used at definition
	fmt.Println("A")
}

func (a *B) Print() {
	fmt.Println("B")
}

func (b *B) Increase() {
	*b = *b + 100

}
