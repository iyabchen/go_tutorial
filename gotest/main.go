package gotest

import "fmt"

func Hello() {
	fmt.Println("hello from the other side")
	hi()
}

func hi() {
	fmt.Println("hi from the same side")
}
