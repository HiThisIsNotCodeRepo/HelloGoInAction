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

