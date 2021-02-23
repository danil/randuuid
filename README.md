randuuid
========

[![Build Status](https://cloud.drone.io/api/badges/danil/randuuid/status.svg)](https://cloud.drone.io/danil/randuuid)
[![Go Reference](https://pkg.go.dev/badge/github.com/danil/randuuid.svg)](https://pkg.go.dev/github.com/danil/randuuid)

Random UUID with seed for Go.  
Source files are distributed under the BSD-style license
found in the [LICENSE](./LICENSE) file.

Install
-------

    go get github.com/danil/randuuid@v0.2.0

Usage
-----

```go
package main

import (
    "fmt"

    "github.com/danil/randuuid"
)

func main() {
    fmt.Println(randuuid.New(42))
    fmt.Println(randuuid.New(42))
}
```

Output:

    538c7f96-b164-4f1b-97bb-9f4bb472e89f <nil>
    538c7f96-b164-4f1b-97bb-9f4bb472e89f <nil>
