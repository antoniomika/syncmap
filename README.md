# syncmap [![Go Reference](https://pkg.go.dev/badge/github.com/antoniomika/syncmap.svg)](https://pkg.go.dev/github.com/antoniomika/syncmap) ![Coveralls](https://img.shields.io/coveralls/github/antoniomika/syncmap)

A generic wrapper around [sync.Map](https://pkg.go.dev/sync#Map).

## How to use

```golang
package main

import (
    "log"

    "github.com/antoniomika/syncmap"
)

func main() {
    testMap := syncmap.New[string, string]()
    testMap.Store("foo", "bar")

    log.Println(testMap.Load("foo"))
}
```
