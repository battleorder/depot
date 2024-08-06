package server

import (
	"net/http"

	"github.com/supabase-community/supabase-go"
)

func NewServer(
  spb *supabase.Client,
) http.Handler {
  mux := http.NewServeMux()

  addRoutes(mux)

  var handler http.Handler = mux
  handler = newLoggingMiddleware(handler)
  handler = newAuthenticatedMiddleware(spb, handler)
  return handler
}
