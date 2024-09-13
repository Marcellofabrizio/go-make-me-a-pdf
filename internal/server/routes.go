package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-make-me-a-pdf/internal/pdf"
	"io"
	"log"
	"net/http"

	"github.com/go-pkgz/routegroup"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := routegroup.Mount(http.NewServeMux(), s.rootUrl)
	r.HandleFunc("GET /health", s.HealthHandler)
	r.HandleFunc("POST /v1/convert-pdf-by-web", s.GeneratePdf)
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

func (s *Server) GeneratePdf(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data map[string]string

	if len(body) > 0 {
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("%+v\n", data)
	}

	orderId := data["orderId"]
	emissionDate := data["emissionDate"]
	licensePlate := data["licensePlate"]

	if orderId == "" {
		err := errors.New("missing orderId")
		log.Fatalf("error handling JSON marshal. Err: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if emissionDate == "" {
		err := errors.New("missing emissionDate")
		log.Fatalf("error handling JSON marshal. Err: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if licensePlate == "" {
		err := errors.New("missing licensePlate")
		log.Fatalf("error handling JSON marshal. Err: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pdf.GeneratePdf(orderId, emissionDate, licensePlate)
}
