package main

import (
    "fmt"
    "net/http"
    "log"
    "math/rand"
    _ "net/http/pprof"
)

// genAuthGUID returns a string attened to an int between 1 and 80000
func genAuthGUID(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "LoadUser%d", rand.Intn(79999)+1)
}

func main() {

    // to use please refer to https://golang.org/pkg/net/http/pprof/
    go func() {
    	log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    http.HandleFunc("/genauthguid", genAuthGUID)

    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
