package main

import (
    "flag"
    // "fmt"
    "log"
    "net/http"

    "github.com/carlmjohnson/feed2json"
)

func main() {
    flag.Parse()
    http.Handle("/api/feed", feed2json.Handler(
        feed2json.StaticURLInjector("https://news.ycombinator.com/rss"), nil, nil, nil))
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

}