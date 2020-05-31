package app

import (
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer(router *mux.Router) *Server {
	s := Server{Router: router}
	return &s
}
