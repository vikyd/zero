# go-zero-check
Golang [The zero value](https://golang.org/ref/spec#The_zero_value)

Check if a variable is its type's `zero value` or `default value`

The variable type can be any type:

- nil
- int
- float
- string
- slice
- struct
- channel
- map
- byte
- interface

# Install

`go get -u github.com/vikyd/go-zero-check`

# Usage

```go
package main
import (
    "fmt"
    "github.com/vikyd/go-zero-check"
)

func main() {
    var v int
    fmt.Println(zero.IsZeroVal(v))
}
```

# Thanks

[newacct](https://stackoverflow.com/a/13906031/2752670)
