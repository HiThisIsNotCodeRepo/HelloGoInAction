package main

import "fmt"

type myType map[string]string

func (m myType) change() {
	m["Hello"] = "World"
}
func main() {
	m := myType{"Hello": ""}
	fmt.Println(m)
	m.change()
	fmt.Println(m)
}
