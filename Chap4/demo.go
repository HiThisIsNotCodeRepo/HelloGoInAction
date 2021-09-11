package main

import "fmt"

func main() {
	sliceDemo1()
	sliceDemo2()
	sliceDemo3()
	sliceDemo4()
	sliceDemo5()
	sliceDemo6()
	sliceDemo7()
}

func sliceDemo1() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice[0] = 25
	fmt.Printf("slice:%v\n", slice)
	fmt.Printf("newSlice:%v\n", newSlice)
}
func sliceDemo2() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Printf("slice:%v\n", slice)
	fmt.Printf("newSlice:%v\n", newSlice)
}
func sliceDemo3() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice := source[2:3:3]
	slice = append(slice, "KiWi")
	fmt.Printf("source:%v\n", source)
	fmt.Printf("slice:%v\n", slice)
}
func sliceDemo4() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice := source[2:3]
	slice = append(slice, "KiWi")
	fmt.Printf("source:%v\n", source)
	fmt.Printf("slice:%v\n", slice)
}
func sliceDemo5() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}

	fmt.Printf("%v\n", append(s1, s2...))
}
func sliceDemo6() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
}
func sliceDemo7() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
}
