# `type-info`
_A library for discovering type information in Go_

## Why
I wanted a small library with simple access to the commonly used parts of the Go `reflect` library.

## Usage
```go
package main

import oi "github.com/nicklarsennz/go-type-info"

func main() {
    // Given a struct `S`
    s_fields, err := oi.StructFields(&S{})
}
```

## Todo:
- `StructMethods`
- `InterfaceMethods`
- `isPublic bool`
- `isPrivate bool`