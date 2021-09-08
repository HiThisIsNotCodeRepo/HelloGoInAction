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

*`var` and `:=``

- When to initialise zero value variable use `var`.
- When to initialise non-zero value or use function returns to initialise use `:=`

*Reference type*

1. `channel`
2. `map`
3. `slice`

*Program termination point*

When return from `main()`, the whole program terminates and in the same time end all goroutine.

*`sync.WaitGroup`*

Suggest use `WaitGroup` to track if goroutine complete the job.

1. Initialise and set `WaitGroup` to the quantity of goroutine.

```
var waitGroup sync.WaitGroup
waitGroup.Add(len(feeds))
```

2. When goroutine finish task decrease `WaitGroup` counter.
3. When `WaitGroup` counter reaches 0 meaning all goroutine exits.

*Key word `range`*

`range` can be used to iterate:

1. Array.
2. String.
3. Slice.
4. Map.
5. Channel.

*`_` in `:=`*

It's actually a placeholder, to ignore the that value.

*map returns*

map returns one value or two values. The differences are:

1. If returns one value it's the search value.
2. If returns two values first one is the search value, second is to bool value to indicate if that key exists in the
   map.

*goroutine*

goroutine is the function runs parallel with main and use key word `go` to start it.

*closure in goroutine*

goroutine anonymous function have 2 ways to access the variable:

1. closure it's suitable for the value not change.
2. arguments in function it's suitable for the value keep changing.

*struct tag*

In struct the tag describe JSON encoding metadata it maps the field of struct to json field.

```
type Feed struct {
   Name string `json:"site"`
   URI string `json:"link"`
   Type string `json:"type"`
}
```

*defer*

The function behind `defer` will always be executed after function return even when exception occurs.

*`interface{}`*

`interface{}` is a special type used with reflect package.

*interface*

interface defines an action that struct or named type needs to implement.

By convention if an interface has only one method and that interface's name should end with er.

*empty struct*

An empty struct consumes zero space.

```
type defaultMatcher struct{}
```

*method receiver*

When declare a receiver for a method, it means that method has been bind with that type.

The receiver could be a value or a pointer to type.

That type's pointer or value can be used to call method.

When needs to mutate the value of the struct set method receiver to pointer.

When using pointer receiver to implements an interface, the interface variable must be a pointer so that interface
variable can call interface method.

When using value receiver to implements an interface, the interface variable can be either pointer or value to call
interface method.

The interface type can receive either value or pointer.

*for range channel*

When channel is closed `for range channel` will exit.

```
// Display 从每个单独的goroutine 接收到结果后
// 在终端窗口输出
func Display(results chan *Result) {
   // 通道会一直被阻塞，直到有结果写入
   // 一旦通道被关闭，for 循环就会终止
   for result := range results {
   fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
      }
}
```