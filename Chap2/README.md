# Chap 2

*Structure of program*

1. Main goroutine
2. Search goroutine
3. Trace goroutine

*Steps*

1. Main goroutine receive data.
2. Main goroutine create search go routine to search.
3. Search go routine will send data back to main goroutine to display.
4. When search go routine complete job it will send message to trace goroutine to notify.
5. Trace goroutine will keep an internal counter when all search go routine complete it will send end signal to main go
   routine.

*Structure*

```
- sample
   - data
      data.json 
   - matchers
      rss.go
   - search
      default.go
      feed.go
      match.go
      search.go
   main.go
```

*main package*

`main()` exists in package `main` is used as program entry.

*package rule*

Every file in the same folder must use the same package name and by convention package name and folder name are the
same.

*`_` in import*

It makes go to do initialisation in package but not to use it.

*How go compiler looks for package*

1. `GOROOT`
2. `GOPATH`

*Uppercase and lowercase*

In package any uppercase identifier will become public, any lowercase will become private, but can access it through
public function indirectly.

*map*

1. Map is a reference type, meaning the value it stores point to the underlying object.
2. Use `make` can initialise a map `var matchers = make(map[string]Matcher)`, without init the map default value
   is `nil`

*function*

In go the key word `func` is used to declare function. The function can have multiple return for example value and
error. Error is the exception to the function, but more graceful, it makes exception just like a value, so we can have
lesser code to handle exception.

Use `:=` to declare and assign has no difference with keyword `var`