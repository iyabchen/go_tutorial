// print.go

package main

import (
	alias "fmt"
)

const PI = 3.14

var name = "this is a name"

type newType int

type gopher struct{}

type golang interface{}

type (
	byte     int8
	rune     int32
	ByteSize int64
)

func main() {

	alias.Println("hello world")
	var a float32 = 100.1
	alias.Println(a)
	b := int(a)
	alias.Println(b)

	var c int = 65
	d := string(c)
	alias.Println(d)
}
