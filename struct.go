package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type human struct {
	Sex int
}

type teacher struct {
	human // if anonymus, then the type was used as name
	Name  string
	Age   int
}

type stA struct {
	stB
	Name string
}

type stB struct {
	Name string
}

func main() {
	a := person{}
	a.Age = 10
	a.Name = "Joe"
	fmt.Println(a)

	b := person{
		Name: "JoeB",
		Age:  15, // must have comma
	}
	fmt.Println(b)

	A(a)
	fmt.Println(a)

	B(&a)
	fmt.Println(a)

	c := &person{
		Name: "JoeC",
		Age:  19,
	}
	fmt.Println(c)

	c.Name = "OK" // no need to do *c.Name
	fmt.Println(c)

	anony := &struct { // anonoymous structure
		Name    string
		Age     int
		Contact struct { // embeded struct
			Phone, City string
		}
	}{
		Name: "anonymous1",
		Age:  111,
	}
	anony.Contact.Phone = "343" // cannot be init with Name, and Age
	anony.Contact.City = "aaa"
	fmt.Println(anony)

	teacher1 := teacher{Name: "abc", Age: 19, human: human{Sex: 0}}
	teacher1.Sex = 1 // embeded into the struct
	teacher1.human.Sex = 1
	fmt.Println(teacher1)

	m := stA{Name: "A", stB: stB{Name: "B"}}
	fmt.Println(m.Name, m.stB.Name) // same name member
}

func A(per person) { // passed the object copy
	per.Age = 111 // cannot modify the object
	fmt.Println("func A", per)

	tmp := per
	tmp.Name = "hans"
	fmt.Println(per)
}

func B(per *person) {
	per.Age = 123
	fmt.Println("func B", per)
}
