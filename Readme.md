
[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

### chai

A tiny wrapper around below packages.

* [google/go-querystring](https://github.com/google/go-querystring)
* [justinas/alice](https://github.com/justinas/alice) Painless middleware chaining for Go
* [go-zoo/bone](https://github.com/go-zoo/bone) A lightweight and lightning fast HTTP Multiplexer for Go

### Example

```go
package main

import "github.com/coderhaoxin/chai"
import "net/http"
import "fmt"

func main() {
  app := chai.New()

  app.Get("/foo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "foo")
  }))

  app.Get("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "bar")
  })

  app.Listen(":3000")
}
```

### License
MIT

[travis-img]: https://img.shields.io/travis/coderhaoxin/chai.svg?style=flat-square
[travis-url]: https://travis-ci.org/coderhaoxin/chai
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: http://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/coderhaoxin/chai
