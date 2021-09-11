package main

import (
	"fmt"
	"unsafe"
)

func main() {
	sliceDemo1()
	sliceDemo2()
	sliceDemo3()
	sliceDemo4()
	sliceDemo5()
	sliceDemo6()
	sliceDemo7()
	sliceDemo8()
	mapDemo3()
	mapDemo4()
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
func sliceDemo8() {
	intArr := [4]int{10, 20, 30, 40}
	intSlice := []int{10, 20, 30, 40}
	strArr := [3]string{"a", "b", "c"}
	strSlice := []string{"a", "b", "c"}
	fmt.Printf("intArr size:%v\n", unsafe.Sizeof(intArr))
	fmt.Printf("intSlice size:%v\n", unsafe.Sizeof(intSlice))
	fmt.Printf("strArr size:%v\n", unsafe.Sizeof(strArr))
	fmt.Printf("strSlice size:%v\n", unsafe.Sizeof(strSlice))
}
func mapDemo1() {
	// ok
	//_ = map[[]string]int{}
	// error
	_ = map[int][]string{}

}

func mapDemo2() {
	var colors map[string]string
	colors["Red"] = "#da1337"
	//panic: assignment to entry in nil map
}

func mapDemo3() {
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}
func mapDemo4() {
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	delete(colors, "Coral")
	delete(colors, "Not exist")
	fmt.Println("After delete")
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}
