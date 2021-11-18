package server

import (
	"github.com/gorilla/mux"
	"github.com/maxheckel/censys-assessment/internal/handlers"
)

func (s *Server) NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", s.Handlers.Healthcheck)
	spa := handlers.SPAHandler{StaticPath: "static/build", IndexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	return r

}
