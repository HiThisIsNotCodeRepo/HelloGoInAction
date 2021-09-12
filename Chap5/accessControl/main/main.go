package main

import (
	"fmt"
	"github.com/qinchenfeng/HelloGoInAction/Chap5/accessControl/entities"
)

func main() {
	a := entities.Admin{Rights: 10}
	a.Name = "Bill"
	a.Email = "bill@email.com"
	fmt.Printf("User: %v\n", a)
}
