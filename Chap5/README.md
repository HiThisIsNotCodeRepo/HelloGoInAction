# Chap5

*Definition of static type language*

In compiling compiler needs to know every value's type.

*`struct`*

`struct` can be used to create a user type

```
type user struct{
    name        string
    email       string
    ext         int
    privileged  bool
}
var bill user
```

*`var`*

By convention, using `var` to init a variable to zero value.

*Create a struct*

Method 1

```
_ = user{
    name:       "Lisa",
    email:      "lisa@email.com",
    ext:        123,
    privileged: true,
}
```

Method 2

```
_ = user{"Lisa", "lisa@email.com", 123, true}
```

*`struct` in `struct`*

Instead of build in type struct also can be embed in struct.

```
type admin struct {
	person user
	level  string
}
```

*Create a type*

Instead of using a struct to create a type, we also can use build in type to create a new type.

```
type Duration int64
```