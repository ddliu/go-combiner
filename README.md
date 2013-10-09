# go-combiner

[![Build Status](https://travis-ci.org/ddliu/go-combiner.png?branch=master)](https://travis-ci.org/ddliu/go-combiner)

`go-combiner` is a golang package to combine strings

## Documentation

[View doc on godoc.org](http://godoc.org/github.com/ddliu/go-combiner)

## Installation

```bash
go get github.com/ddliu/go-combiner
```

## Usage

```go
package main

import (
    "github.com/ddliu/go-combiner"
    "fmt"
)

func main() {
    c := NewCombiner()

    // Add a piece of const(got go_)
    c.Add("const", map[string]interface{} {
        "const": "go_",
    })

    // Add a piece of chars(got aaa, aab, aac, aba, abb, abc...)
    c.Add("chars", map[string]interface{} {
        "chars": "ab",
        "maxlength": 3,
    })

    // Add a piece of choice(got Mac, Linux, Windows)
    c.Add("choice", map[string]interface{} {
        "choices": []string{"Mac", "Linux", "Windows"},
    })

    // Add a piece of dict word(got words start with letter a)
    c.Add("dict", map[string]interface{} {
        "dict": "/usr/share/dict/words",
        "pattern": "[a].*",
    })

    // Walk through all the combinations of above pieces
    c.Walk(func(s string) bool {
        fmt.Println(s)
        return true
    })
}
```

## Changelog

### v0.1.0 (2013-10-09)

Initial release