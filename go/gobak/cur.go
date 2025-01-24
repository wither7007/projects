package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/aoliveti/curling"
)

func main() {
    req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Add("If-None-Match", "foo")

    cmd, err := curling.NewFromRequest(req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(cmd)
}
