# syncmap [![Go Reference](https://pkg.go.dev/badge/github.com/antoniomika/syncmap.svg)](https://pkg.go.dev/github.com/antoniomika/syncmap) [![Coverage Status](https://coveralls.io/repos/github/antoniomika/syncmap/badge.svg?branch=main)](https://coveralls.io/github/antoniomika/syncmap?branch=main)

A generic wrapper around [sync.Map](https://pkg.go.dev/sync#Map).

## [How to use](https://play.golang.com/p/9OmDSbzNC5Z)

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
