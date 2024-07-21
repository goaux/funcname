# funcname

`funcname` is a Go package that provides utility functions for retrieving and
manipulating function names. It offers various methods to obtain function names
from different contexts, such as runtime program counters, interface values,
and the current call stack.

## Installation

To install the `funcname` package, use the following command:

```
go get github.com/goaux/funcname
```

## Usage

Here are some examples of how to use the `funcname` package:

```go
package main

import (
    "fmt"
    "github.com/goaux/funcname"
)

func main() {
    // Get the name of the current function
    fmt.Println(funcname.This())

    // Get the name of a function from an interface
    fmt.Println(funcname.Of(fmt.Println))

    // Split a function name into package and function components
    pkg, fn := funcname.Split("github.com/user/pkg.Func")
    fmt.Printf("Package: %s, Function: %s\n", pkg, fn)
}
```

## API

The `funcname` package provides the following main functions:

- `Of(v interface{}) string`: Returns the full name of the function or method represented by `v`.
- `SplitOf(v interface{}) (pkgname, name string)`: Returns the package name and function name of the function or method represented by `v`.
- `This() string`: Returns the name of the calling function.
- `SplitThis() (pkgname, name string)`: Returns the package name and function name of the calling function.
- `Caller(skip int) string`: Returns the name of the calling function, skipping the specified number of stack frames.
- `SplitCaller(skip int) (pkgname, name string)`: Returns the package name and function name of the calling function, skipping the specified number of stack frames.
- `ForPC(pc uintptr) string`: Returns the name of the function corresponding to the given program counter.
- `SplitForPC(pc uintptr) (pkgname, name string)`: Returns the package name and function name corresponding to the given program counter.
- `Split(s string) (pkgname, name string)`: Separates a fully qualified function name into package name and function name.
