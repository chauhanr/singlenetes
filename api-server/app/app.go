package app

import (
	"fmt"
	"log"
	"net/http"

	store "github.com/chauhanr/singlenetes/api-server/store"
	"github.com/gorilla/mux"
)

//Startup methid is used to start the api server
func Startup(port int, r *mux.Router, etcdCtl store.EtcdCtl, hClient *http.Client) {
	s := NewServer(r, etcdCtl)
	s.ApiRoutes()

	// configure and start watcher
	w := NewWatcher(etcdCtl.GetEtcdClient(), hClient)
	go w.Start()
	defer w.Close()

	log.Printf("Starting api-server on port %d\n", port)
	portAddr := fmt.Sprintf(":%d", port)

	http.ListenAndServe(portAddr, s.Router)

}
