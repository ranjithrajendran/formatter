# formatter

[![GoDoc Widget]][GoDoc]

`formatter` is a Golang utility package to make the normal float amount to human readable cash format, by proper `,` (coma) separation.  

## Install

`go get github.com/ranjithrajendran/formatter`

## Functions

### ReadableIndianCash( )

**function signature:** `func ReadableIndianCash(num float64) string`

**functionality:** This program convers the normal float amount to indian cash format.

**example**

```go
package main

import (
	"fmt"
	"github.com/ranjithrajendran/formatter"
)

func main() {
	fmt.Println(formatter.ReadableIndianCash(100))
  // will print 100.00
  
  fmt.Println(formatter.ReadableIndianCash(5462.54673742))
  // will print 5,462.55
  
  fmt.Println(formatter.ReadableIndianCash(345676.63490))
  // will print 3,45,676.63
  
  // max limit
  fmt.Println(formatter.ReadableIndianCash(99999999999999.9899999))
  // will print 99,99,999,99,99,999.98
  
}
```


[GoDoc]: https://godoc.org/github.com/ranjithrajendran/formatter
[GoDoc Widget]: https://godoc.org/github.com/go-chi/chi?status.svg
