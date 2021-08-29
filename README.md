# Quickstart

```go
package main

import (
    "fmt"
    "log"

    "github.com/psotou/fintoc-sdk/fintoc"
)

func main() {
    client, err := fintoc.NewClient("<secret>")
    if err != nil {
        log.Fatal(err)
    }

    link := client.Link.Get("<linkToken>")
    fmt.Prinln(link)
}
```

# TO-DO
                                                                                            
+ [ ] Add query params to movements call
+ [ ] Add methods PATCH and DELETE for link object.
+ [ ] Add tests.

So far, only a limited amount of FINTOC API resources are going to be covered. More resources should be added in the future.