package middleware

import (
    "net/http"
    "sync"
    "time"
)

type RateLimiter struct {
    visitors map[string]*visitor
    mu       sync.Mutex
}

type visitor struct {
    tokens int
    last   time.Time
}

func RateLimiter(next http.Handler) http.Handler {
    rl := &RateLimiter{visitors: make(map[string]*visitor)}
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        rl.mu.Lock()
        v, ok := rl.visitors[ip]
        if !ok {
            v = &visitor{tokens: 100, last: time.Now()}
            rl.visitors[ip] = v
        }
        if v.tokens <= 0 {
            rl.mu.Unlock()
            http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        v.tokens--
        rl.mu.Unlock()
        next.ServeHTTP(w, r)
    })
}
