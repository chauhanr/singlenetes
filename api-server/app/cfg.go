package app

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	ApiServerConfig ApiServerConfig `yaml:"api-server"`
}

type ApiServerConfig struct {
	Port       int        `yaml:"port"`
	EtcdConfig EtcdConfig `yaml:"etcd"`
}

type EtcdConfig struct {
	Endpoints      []string `yaml:"endpoints"`
	ContextTimeout int      `yaml:"context-timeout"`
}

/*validattion function will check for all mandatory functions.
  also wil check for correctness of data
*/
func (s *ServerConfig) validate() []error {
	return nil
}

/*
   Will load the configurations using the file path provided.
   it will error out if either config file is not found or there is a parsing error.
*/
func (s *ServerConfig) load(config string) error {
	// Load the config file
	f, err := os.Open(config)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	err = dec.Decode(s)
	if err != nil {
		return err
	}
	return nil
}

//LoadAndValidateApiConfig method will validate the api configurations.
func (s *ServerConfig) LoadAndValidateApiConfig(path string) error {
	err := s.load(path)
	if err != nil {
		return fmt.Errorf("error loading configurations for api server: %s", err)
	}

	errs := s.validate()
	if errs != nil {
		for i, e := range errs {
			_ = fmt.Errorf("%d. Error - %s", i, e)
		}
		return errors.New("multiple validation errors in config")
	}
	return nil
}
