package main

import (
	"flag"
	"fmt"
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
		fmt.Printf("Error loading/validating server config: %s\n", err)
		return
	}
	etcdConfig := cfg.ApiServerConfig.EtcdConfig

	cli, err := etcdCli(etcdConfig)
	if err != nil {
		fmt.Printf("Error connecting to ectd instance: %s\n", err)
		return
	}
	client := store.EtcdCtlImpl{cli}
	defer cli.Close()

	hClient := http.DefaultClient
	app.Startup(cfg.ApiServerConfig.Port, r, client, hClient)

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
