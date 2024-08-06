package server

import (
	"log/slog"
	"net/http"
	"time"
)

func newLoggingMiddleware(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    now := time.Now()

    h.ServeHTTP(w, r)

    slog.Info(
      "",
      "method", r.Method,
      "path", r.URL.EscapedPath(),
      "host", r.Host,
      "fwd", r.Header.Get("X-Forwarded-For"),
      // "status_code", r.Response.StatusCode,
      "ns", time.Now().Sub(now).Nanoseconds(),
    )
  })
}
