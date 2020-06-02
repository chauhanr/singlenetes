package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) ApiRoutes() {
	var api = s.Router.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(s.NotFoundHandler())
	s.apiV1Routes(api)
	s.apiV1alphaRoutes(api)
}

func (s *Server) apiV1Routes(subRouter *mux.Router) {
	var api = subRouter.PathPrefix("/v1").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(s.NotFoundHandler())
	api.HandleFunc("/status", s.StatusV1())
}

func (s *Server) apiV1alphaRoutes(subRouter *mux.Router) {
	var api = subRouter.PathPrefix("/v2alpha").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(s.NotFoundHandler())
	api.HandleFunc("/status", s.StatusV1alpha())
}
