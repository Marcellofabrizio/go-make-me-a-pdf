package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.HealthHandler)
	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := make(map[string]string)
	resp["message"] = "Go Make Me a PDF"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := make(map[string]string)
	resp["message"] = "Up and Running!"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
