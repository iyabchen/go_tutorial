// strconv.go

package main

import (
	alias "fmt"
	"strconv"
)

const PI = 3.14

func main() {

	var c int = 65
	d := string(c)
	alias.Println(d)
	e := strconv.Itoa(c)
	alias.Println(e)
	f, _ := strconv.Atoi(e)
	alias.Println(f)
}
