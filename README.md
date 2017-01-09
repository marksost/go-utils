# go-utils

A collection of useful Golang utility methods and libraries

---

[![GoDoc](https://godoc.org/github.com/marksost/go-utils?status.svg)](http://godoc.org/github.com/marksost/go-utils) [![Build Status](https://travis-ci.org/marksost/go-utils.svg?branch=master)](https://travis-ci.org/marksost/go-utils)

## Installation

Install:

```go
go get github.com/marksost/go-utils
```

Import:

```go
import (
	goutils "github.com/marksost/go-utils"
)
```

## Documentation:

https://godoc.org/github.com/marksost/go-utils

## Usage

```go
b := true
fmt.Printf("Type: %T, Value:%v\n", b, b) // Type: bool, Value:true
s := goutils.Bool2String(b)
fmt.Printf("Type: %T, Value:%v\n", s, s) // Type: string, Value:true
```
