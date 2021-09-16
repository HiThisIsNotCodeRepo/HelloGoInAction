# Chap8

*go std lib*

[Standard library](https://pkg.go.dev/std)

*$GOROOT/src*

Source code.

*$GOROOT/pkg*

Pre compiled file used to build program.

*log is thread safe*

*Log configuration method 1*

Use `log` method.

```
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
```

*Log configuration method 2*

Use `*log.Logger` value.

```
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
```

*Usage of tag in `struct` when used in json decoding*

It provides the metadata of every field, map json to struct field.

*The types json can be decoded*

1. Struct.
2. Map `map[string]interface{}`

