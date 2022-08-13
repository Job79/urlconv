Package urlconv is a small helper that maps url.Values into a given struct.

Example:
```go
type MyStruct struct {
	Name string `url:"name"`
	Age  int    `url:"age"`
}

values := url.Values{
	"name": []string{"John"},
	"age":  []string{"42"},
}

var s MyStruct
urlconv.Unmarshal(values, &s)
fmt.Println(s.Name, s.Age)
// Output: John 42
```

The struct tag "url" is used to map the url.Values to the struct
and is required. Only public structs and fields can be mapped.

The following types are supported:
- string
- int
- []string
- bool
