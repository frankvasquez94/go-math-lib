# go-math-lib

This is an example test of how to create a library in go.

## Steps to create a go library

1. Create a repository on github, in this case go-math-lab.
2. Clone the repository, in this case using the path github.com/username/
3. Create the module
```
go mod init github.com/frankvasquez94/go-math-lib
```
4. Write the code of the library.
5. push the code to your repository.
```
git puh origin main
```
6. Create a tag
```
git tag v0.0.1
```
7. push the code to the tag.
```
git push origin v0.0.1
```

At this point our library is ready!!


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

