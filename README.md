# go-clipfile
> Extract file paths from macOS clipboard in Go language.

## installation
```sh
go get -u github.com/afeiship/go-clipfile
```

## usage
```go
package main

import (
	"fmt"
	"github.com/afeiship/go-clipfile"
)

func main() {
	path := clipfile.GetPath()
	if path == "" {
		fmt.Println("Failed to get file path")
	}
	fmt.Println(path)
}
```