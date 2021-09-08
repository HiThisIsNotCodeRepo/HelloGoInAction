# Chap 3

*package*

Package is a group of `.go` files. By convention package name is same as folder name , lowercase letter.

Because the package is imported by full path, so different package can have the same package name.

Also, after package import, the imported package name can be changed.

go will compile `main` package into binary file.

Function `main()` exists in `main` package , it's the entry of the program.

*How go compiler import package*

Suppose Go is installed in `/usr/local/go` and `GOPATH` is set as `/home/myproject:/home/mylibraries/`, go compiler will
look for `net/http` in below order:

1. `/usr/local/go/src/pkg/net/http`
2. `/home/myproject/src/net/http`
3. `/home/mylibraries/src/net/http`

*How go compiler import remote package*

```
import "github.com/spf13/viper"
```

go compiler will start searching from `GOPATH` , if it doesn't exist use `go get` to download the package.

*Rename import*

```
import (
    "fmt"
    myfmt "mylib/fmt"
)
```

*init()*

One package can have any number of `init()`, they will be executed before program starts.

*go build*

Build all package under `chapter3`

```
go build github.com/goinaction/code/chapter3/...
```

*go run*

Build and run.

```
go run wordcount.go
```

*go vet*

`go vet` catches error include:

1. `Printf` the placeholder not match the type.
2. Function signature error.
3. Wrong struct tag.
4. When init struct with literal value , not specify field.

*go fmt*

`go fmt` will format code to source code style.

*go doc vs godoc*

Use `go doc` command to view documentation of specific package in terminal.

```
go doc tar
```

Run `godoc` to start a HTTP documentation server.

```
godoc -http=:6060
```

*Proper way to doc/comment*

1. Add comment right above package,function,type and global variable.
2. If it needs to write a large documentation for package can write in a separate `doc.go` file, just above package
   name.

*How to write a sharable code*

1. Store package in `GOPATH/src`.
2. Make package small.
3. Use `go fmt` to format code.
4. Write more comment.

*What is the purpose of go mod vendor?*

go `imports` neither care dependency version nor dependency is existed.

`mod` takes care of version information.

`vendor` download dependency so when build the program we won't worry that dependency doesn't exist.