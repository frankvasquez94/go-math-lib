# go-math-lib

## install

```http request
go get github.com/frankvasquez94/go-math-lib@v0.0.2
```

## Usage

```go
package main

import (
	"fmt"
	mathlb "github.com/frankvasquez94/go-math-lib"
)

func main() {
	p := fmt.Println
	sp := fmt.Sprintf
	sum := mathlb.AddInt(3, 4)
	p(sp("%d + %d = %d", 3, 4, sum))

	minus := mathlb.MinusInt(7, 4)
	p(sp("%d - %d = %d", 7, 4, minus))

	multiply := mathlb.MultiplyInt(4, 4)
	p(sp("%d * %d = %d", 4, 4, multiply))

	division := mathlb.DivideInt(16, 4)
	p(sp("%d / %d = %d", 16, 4, division))

	module := mathlb.Module(7, 5)
	p(sp("%d module %d = %d", 7, 5, module))
}


```


## Functions

AddInt(a, b int) int

MinusInt(a, b int) int

MultiplyInt(a, b int) int

DivideInt(a, b int) int

Module(a, b int) int