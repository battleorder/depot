package server

import "net/http"

func addRoutes(mux *http.ServeMux) {
	mux.Handle("GET /v1/units", handleGetUnits())
}
