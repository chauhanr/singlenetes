package main

import (
	"github.com/chauhanr/golang-web/10-api-server/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	s := app.NewServer(r)
	s.ApiRoutes()

	log.Println("Starting api-server on port 9991")
	http.ListenAndServe(":9991", s.Router)
}
