package server

import (
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/maxheckel/censys-assessment/internal/handlers"
)


func (s *Server) NewRouter() *mux.Router{
	r := mux.NewRouter()
	//// Setup common middleware
	cors := ghandlers.CORS(
		ghandlers.AllowedMethods([]string{"GET", "HEAD", "PUT", "POST", "OPTIONS", "PATCH", "DELETE"}),
		ghandlers.AllowedHeaders([]string{
			"Accept", "Accept-Language", "Content-Language", "Origin",
			"X-Requested-With", "Content-Type", "Authorization",
			// Supports the Newrelic distributed tracing headers
			// to link browser transactions
			"Newrelic", "traceparent", "tracestate",
		}),
		ghandlers.AllowedOrigins([]string{"*"}),
	)
	r.Use(cors)
	r.HandleFunc("/healthcheck", s.Handlers.Healthcheck)
	r.HandleFunc("/ip/{address}", s.Handlers.GetIPDetails)

	spa := handlers.SPAHandler{StaticPath: "static/build", IndexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	return r

}
