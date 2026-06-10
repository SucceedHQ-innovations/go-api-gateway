package main

import (
    "log"
    "net/http"
    "os"
    "github.com/SucceedHQ-innovations/go-api-gateway/internal/server"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    srv := server.New()
    log.Printf("API Gateway listening on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, srv.Router()))
}
