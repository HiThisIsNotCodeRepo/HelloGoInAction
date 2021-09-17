# Chap9

*Unit test*

1. Basic test
2. Table test
3. Mock test

*When will unit test be considered as fail*

When test doesn't call `t.Fatal` or `t.Error`

*Purpose of mock test*

It's not good for your test depends on external resource, because external resource is not guaranteed and without it
your can't run your test as your test will fail definitely.

*Snippet to test endpoint*

```
req, err := http.NewRequest("GET", "/sendjson", nil)
rw := httptest.NewRecorder()
http.DefaultServeMux.ServeHTTP(rw, req)
```

*Write example for package*

The function name must in format of `Example+FunctionName`

*Init a godoc server*

```
godoc -http=":3000"
```

*Command run benchmark test*

```
go test -v -run="none" -bench="." -benchtime="3s"
go test -v -run="none" -bench="." -benchtime="3s" -benchmem
```

*Reset timer*

Put `b.ResetTimer()` just before `for` to make result more accurate.
