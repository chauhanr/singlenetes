package main

import (
	"flag"
	"fmt"
	"github.com/chauhanr/singlenetes/api-server/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var config string

func init() {
	flag.StringVar(&config, "c", "", "api-configuration")
}

func main() {
	flag.Parse()
	r := mux.NewRouter()
	if config == "" {
		fmt.Printf("Error : api configurations is mandatory for the api server to start.\n")
		fmt.Printf("Usage: api-server -c /path/to/config.yml\n")
		return
	}
	cfg := app.ServerConfig{}
	err := cfg.LoadAndValidateApiConfig(config)
	port := cfg.ApiServerConfig.Port

	if err != nil {
		fmt.Errorf("Error loading/validating server config: %s\n", err)
		return
	}

	s := app.NewServer(r)
	s.ApiRoutes()

	log.Printf("Starting api-server on port %d\n", port)
	portAddr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(portAddr, s.Router)
}
