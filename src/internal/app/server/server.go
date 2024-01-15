package server

import (
	"net/http"

	"github.com/clyentsoftwares/currency-api/src/internal/app/currency"
	"github.com/gorilla/mux"
)

type Server struct {
	router          *mux.Router
	currencyService *currency.Service
}

func NewServer() *Server {
	s := &Server{
		router:          mux.NewRouter(),
		currencyService: currency.NewService(),
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
