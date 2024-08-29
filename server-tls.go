package main

import (
    "fmt"
    "net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "%v: %v\n", "Host", req.Host)
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    http.HandleFunc("/", headers)
    http.ListenAndServe(":80", nil)
    http.ListenAndServeTLS(":443", "server_certificate.pem", "private_key.pem", nil)
}

