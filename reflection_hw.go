package main

import (
	"fmt"
	"reflect"
)

type AnimalBehavior interface {
	Shout(content string, times int, receiver string) []int
}

type Pet struct {
	Id   int
	Age  int
	Name string
}

type Cat struct {
	Pet
	cat_type string
}

type Dog struct {
	Pet
	dog_type string
}

func (cat Cat) Shout(content string, times int, receiver string) (ret []int) {
	fmt.Println("miao~ ", content, "shouted ", times, " times a day, calling its ", receiver)
	ret = []int{100, 300}
	return ret
}

func (dog Dog) Shout(content string, times int, receiver string) []int {
	fmt.Println("wang~ ", content, "shouted ", times, " times a day, calling its ", receiver)
	return []int{200, 400}
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if t.Kind() != reflect.Struct {
		fmt.Println("info getting struct error")
		return
	}

	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, v.Interface())
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}

func main() {
	peta := Dog{Pet: Pet{1, 3, "bobo"}, dog_type: "husky"}
	petb := Cat{Pet: Pet{2, 4, "mimi"}, cat_type: "garfield"}
	return_peta := peta.Shout("hi", 2, "master")
	fmt.Println("return peta", return_peta)
	petb.Shout("hello", 1, "sleep")

	Info(peta)

	v1 := reflect.ValueOf(petb)
	m1 := v1.MethodByName("Shout")
	if m1.IsValid() {
		args := []reflect.Value{reflect.ValueOf("no content"), reflect.ValueOf(1), reflect.ValueOf("master")}

		rets := m1.Call(args) // returned []reflect.Value
		fmt.Println(rets[0])
	}

}
