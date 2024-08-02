package server

import (
	"net/http"

	"github.com/go-kit/log"
)

func addRoutes(mux *http.ServeMux, lgr log.Logger) {
  mux.Handle("GET /v1/units", handleGetUnits(lgr))
}
