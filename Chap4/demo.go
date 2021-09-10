package main

import "fmt"

func main() {
	sliceDemon1()
	sliceDemon2()
}

func sliceDemon1() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice[0] = 25
	fmt.Printf("slice:%v\n", slice)
	fmt.Printf("newSlice:%v\n", newSlice)
}
func sliceDemon2() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Printf("slice:%v\n", slice)
	fmt.Printf("newSlice:%v\n", newSlice)
}
