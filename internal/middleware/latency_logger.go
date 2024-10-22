package middleware

import (
    "fmt"
    "net/http"
    "time"
)

func LatencyLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()
        next.ServeHTTP(w, r)
        latency := time.Since(startTime)
        fmt.Printf("Request Latency: %v\n", latency)
    })
}
