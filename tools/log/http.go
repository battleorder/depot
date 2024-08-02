package log

import (
	"net/http"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func NewLoggingMiddleware(logger log.Logger, h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    defer level.Info(logger).Log(
      "msg", "request received",
      "method", r.Method,
      "path", r.URL.Path,
    )
    h.ServeHTTP(w, r)
  })
}
