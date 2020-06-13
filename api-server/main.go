package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chauhanr/singlenetes/api-server/app"
	"github.com/chauhanr/singlenetes/api-server/store"
	"github.com/gorilla/mux"
	"go.etcd.io/etcd/clientv3"
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
	if err != nil {
		fmt.Errorf("Error loading/validating server config: %s\n", err)
		return
	}
	etcdConfig := cfg.ApiServerConfig.EtcdConfig

	cli, err := etcdCli(etcdConfig)
	if err != nil {
		fmt.Printf("Error connecting to ectd instance: %s\n", err)
		return
	}
	client := store.EtcdCtl{cli}
	defer cli.Close()

	s := app.NewServer(r, &client)
	s.ApiRoutes()

	// configure and start watcher
	hClient := http.DefaultClient
	w := app.NewWatcher(&client, hClient)
	go w.Start()
	defer w.Close()

	port := cfg.ApiServerConfig.Port
	log.Printf("Starting api-server on port %d\n", port)
	portAddr := fmt.Sprintf(":%d", port)

	http.ListenAndServe(portAddr, s.Router)

}

func etcdCli(cfg app.EtcdConfig) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: time.Duration(cfg.ContextTimeout) * time.Second,
	})
	if err != nil {
		return cli, err
	}
	return cli, nil
}
