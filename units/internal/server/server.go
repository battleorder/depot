package server

import (
	"net/http"

	"github.com/go-kit/log"
	toolslog "github.com/battleorder/depot/tools/log"
	"github.com/supabase-community/supabase-go"
)

func NewServer(
  logger log.Logger,
  spb *supabase.Client,
) http.Handler {
  mux := http.NewServeMux()

  addRoutes(mux, logger)

  var handler http.Handler = mux
  handler = toolslog.NewLoggingMiddleware(logger, handler)
  handler = newAuthenticatedMiddleware(logger, spb, handler)
  return handler
}
