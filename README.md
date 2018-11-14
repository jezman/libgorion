# libgorion

[![Build Status](https://travis-ci.org/jezman/libgorion.svg?branch=master)](https://travis-ci.org/jezman/libgorion) [![codecov](https://codecov.io/gh/jezman/libgorion/branch/master/graph/badge.svg)](https://codecov.io/gh/jezman/libgorion) [![Go Report Card](https://goreportcard.com/badge/github.com/jezman/libgorion)](https://goreportcard.com/report/github.com/jezman/libgorion)

Library for access control system NVP Bolid "Orion Pro"

## Examples

- companies list

```go
package main

import (
    "fmt"
    "os"

    "github.com/jezman/libgorion"
    _ "github.com/denisenkom/go-mssqldb"
)

func main() {
    dsn := os.Getenv("BOLID_DSN")
    db, err := libgorion.OpenDB(dsn)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    companies, err := db.Company()
    if err != nil {
        fmt.Println(err)
    }

    for num, company := range companies {
        fmt.Println(num, company.Name, company.WorkersCount)
    }
}
```

## More examples

- [gorion](https://github.com/jezman/gorion) based on this library

## License

MIT © 2018 jezman
