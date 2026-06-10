package server

import (
    "net/http"
    "github.com/SucceedHQ-innovations/go-api-gateway/internal/middleware"
)

type Server struct {
    router *http.ServeMux
}

func New() *Server {
    s := &Server{router: http.NewServeMux()}
    s.registerRoutes()
    return s
}

func (s *Server) Router() http.Handler {
    return middleware.RateLimiter(middleware.Logger(s.router))
}

func (s *Server) registerRoutes() {
    s.router.HandleFunc("/api/v1/health", s.handleHealth)
    s.router.HandleFunc("/api/v1/proxy/", s.handleProxy)
}
