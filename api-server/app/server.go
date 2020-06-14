package app

import (
	store "github.com/chauhanr/singlenetes/api-server/store"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	cli    store.EtcdCtl
}

func NewServer(router *mux.Router, ctl store.EtcdCtl) *Server {
	s := Server{Router: router, cli: ctl}
	return &s
}
