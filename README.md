# Quickstart

```go
package main

import (
    "fmt"
    "log"

    "github.com/psotou/fintoc-sdk/fintoc"
)

func main() {
    client, err := fintoc.NewClient("secret")
    if err != nil {
        log.Fatal(err)
    }

    link := client.Link.Get("linkToken")
    fmt.Prinln(link)
}
```

# Disclaimer

This ~yet another carbon copy~ library is heavily influenced by the Python SDK devoleped by the guys at [Fintoc](https://fintoc.com/).

The Go library is in a very early stage in which you can only used certain endpoints, namely, the **link**, **account** and **movements** endpoints.

The whole API, with all its endpoints, should be wrapped up soon, though.

# HOW-TO

## The `Link` interface

The `Link` interface comes with the methods `All` and `Get` (very soon it'll come with the `Update`and `Delete` methods as well).

The available methods can be used as follows:

```go
client, err := fintoc.NewClient("secret")
if err != nil {
    log.Fatal(err)
}

link := client.Link.Get("linkToken") // to return one link object
links := client.Link.All()           // to return all of them
for _, l := range links {
    fmt.Println(l.Id)
}
```

## The `Account` interface

Similarly, the `Account` interface comes with the methods `All` and `Get`, which can be used as:

```go
account := link.Account.Get("accountId")
accounts := link.Account.All()
for _, acc := range accounts {
    fmt.Println(acc.Id)
}
```

## The `Movement` interface

The `Movement` interface also comes with the methods `All` and `Get`:

```go
movement := account.Movement.Get("movementId")
movements := account.Movement.All()
for _, mov := range movements {
    fmt.Println(mov.Id)
}
```

Howeverm, the `All` method of this interface allows for the use query params like this:

```go
params := fintoc.Params{Since: "2021-08-01", Until: "2021-08-31", PerPage: "100"}
movements := account.Movement.All(params)
for _, mov := range movements {
    fmt.Println(mov.Id)
}
```

# TO-DO
                                                                                            
+ [x] Add query params to movements call
+ [ ] Add methods PATCH and DELETE for link object.
+ [ ] Add tests.

So far, only a limited amount of FINTOC API resources are going to be covered. More resources should be added in the future.