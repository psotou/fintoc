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

This ~yet another carbon copy~ library is heavily influenced by the [Python SDK](https://github.com/fintoc-com/fintoc-python) devoleped by the guys at [Fintoc](https://fintoc.com/).

This Go library is in a very early stage in which you can only used certain endpoints, namely, the **link**, **account** and **movements** endpoints. Though, the whole API, with all its endpoints, should be wrapped up soon.

# HOW-TO

## The `Link` interface

The `Link` interface comes with the methods `All`, `Get`, `Update`and `Delete`.

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

To update a link (that is, changing its active status to either true or false), we do as follows:

```go
link := client.Link.Get("linkToken") // to return one link object
link.Update(false)                   // this will print the http status code of the request
```

To delete a link, we should first look for the `linkId` by running the `client.Link.All()` method. Once we get the id, we do:

```go
client.Link.Delete("linkId")        // will also print the http status code of the request
```

Bear in mind that both the `Update`and `Delete` are methods that do not return a value, but rather perform an action.

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

However, the `All` method of this interface allows for the use query params like the following:

```go
params := fintoc.Params{Since: "2021-08-01", Until: "2021-08-31", PerPage: "100"}
movements := account.Movement.All(params)
for _, mov := range movements {
    fmt.Println(mov.Id)
}
```

# TO-DO
                                                                                            
+ [x] Add query params to movements call
+ [x] Add methods PATCH and DELETE for link object
+ [ ] Add tests

So far, only a limited amount of FINTOC API resources are going to be covered. More resources should be added in the future.