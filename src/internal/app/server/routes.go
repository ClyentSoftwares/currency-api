package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) routes() {
	s.router.HandleFunc("/convert", s.handleConvert).Methods("GET")
	s.router.HandleFunc("/rate", s.handleRate).Methods("GET")
	s.router.HandleFunc("/rates", s.handleRates).Methods("GET")
}

func (s *Server) handleConvert(w http.ResponseWriter, r *http.Request) {
	amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)
	if err != nil {
		sendError(w, "Invalid amount", http.StatusBadRequest)
		return
	}
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	convertedAmount, err := s.currencyService.Convert(amount, from, to)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, map[string]interface{}{"amount": convertedAmount})
}

func (s *Server) handleRate(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	rate, err := s.currencyService.GetRate(from, to)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, map[string]interface{}{"rate": rate})
}

func (s *Server) handleRates(w http.ResponseWriter, r *http.Request) {
	base := r.URL.Query().Get("base")
	if base == "" {
		base = "USD"
	}

	rates, err := s.currencyService.GetAllRates(base)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, map[string]interface{}{"rates": rates})
}

func sendJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": message})
}
