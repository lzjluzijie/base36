# base36
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/lzjluzijie/base36)
[![Goreportcard](https://goreportcard.com/badge/github.com/lzjluzijie/base36)](https://goreportcard.com/report/github.com/lzjluzijie/base36)

Package base36 implements base36 encoding

# Installation
`go get github.com/lzjluzijie/base36`

# Examples
Encode
```go
data := []byte("Hello World!")
encoded := base36.Encode(data)

fmt.Println(string(encoded))
```

Decode
```go
encoded := "2678lx5gvmsv1dro9b5"
decoded, err := base36.Decode(encoded)

if err != nil{
  fmt.Println(err.Error())
  return
}

fmt.Println(string(decoded))
```
