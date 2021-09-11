package main

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

type Duration int64

func typeDemo1() {
	_ = user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	_ = user{"Lisa", "lisa@email.com", 123, true}
}
