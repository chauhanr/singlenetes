package app

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
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

func (s *ServerConfig) LoadAndValidateApiConfig(path string) error {
	err := s.load(path)
	if err != nil {
		fmt.Errorf("Error loading configurations for api server: %s.\n", err)
		return err
	}

	errs := s.validate()
	if errs != nil {
		for i, e := range errs {
			fmt.Errorf("%d. error - %s\n", i, e)
		}
		return errors.New("multiple validation errors in config")
	}
	return nil
}
