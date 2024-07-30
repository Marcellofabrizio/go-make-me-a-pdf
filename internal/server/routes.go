package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-pkgz/routegroup"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := routegroup.Mount(http.NewServeMux(), s.rootUrl)
	r.HandleFunc("GET /health", s.HealthHandler)
	return r
}

func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Up and Running!"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
