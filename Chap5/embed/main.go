package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name, email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user:  user{name: "john smith", email: "john@yahoo.com"},
		level: "super",
	}
	sendNotification(&ad)
}
