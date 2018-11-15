# libgorion

[![Build Status](https://travis-ci.org/jezman/libgorion.svg?branch=master)](https://travis-ci.org/jezman/libgorion) [![codecov](https://codecov.io/gh/jezman/libgorion/branch/master/graph/badge.svg)](https://codecov.io/gh/jezman/libgorion) [![Go Report Card](https://goreportcard.com/badge/github.com/jezman/libgorion)](https://goreportcard.com/report/github.com/jezman/libgorion)

Small library for access control system NVP Bolid "Orion Pro"

## Features

All features available from `Datastore` interface

```go
AddWorker(string) error
DeleteWorker(string) error
DisableWorkerCard(string) error
EnableWorkerCard(string) error
Company() ([]*Company, error)
Doors() ([]*Door, error)
Workers(string) ([]*Worker, error)
Events(string, string, string, uint, bool) ([]*Event, error)
EventsValues() ([]*Event, error)
EventsTail(time.Duration, string) error
WorkedTime(string, string, string, string) ([]*Event, error)
```

## Examples

Set environment variable `BOLID_DSN`

```sh
export BOLID_DSN="server=127.0.0.1;user id=username;password=passwd;database=base"
```

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
