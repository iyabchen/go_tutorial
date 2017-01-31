package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello{
	fmt.Println("Hello from user", u.Name)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("ERR")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")

	if f.IsValid() && f.Kind() == reflect.String {
		f.SetString("ppp")
	}

}

type Manager struct {
	User
	title string
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct { // if reflect failed to be a struct
		fmt.Println("ERR")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func main() {
	u := User{1, "OK", 12}
	u.Hello()
	Info(u)

	m := Manager{User: User{1, "OK", 33}, title: "manager"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0)) //Anonymous: true
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

	Set(&u)
	fmt.Println(u)

	v1 := reflect.ValueOf(u)
	m1 := v1.MethodByName("Hello")
	fmt.Println(m1.IsValid())
	// args := []reflect.Value{reflect.ValueOf("joe")}
	m1.call(nil)

}
