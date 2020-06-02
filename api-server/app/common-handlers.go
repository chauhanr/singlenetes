package app

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/chauhanr/singlenetes/api-server/scheme"
	"gopkg.in/yaml.v2"
)

/*
   package to store common handlers like Not found handlers or somehting that can be used across all versions of the
   api that we define.
*/

func (s *Server) NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)

	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("response error: %s\n", err)
	}
}

/* This method assumes that the interface or model that enters
   will implement an interface that asks all models to implement a validate function
   the validate function will check for the validity of the model and return error in case
   failure
*/
func decodeJson(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	/* check for validity of the model*/
	if validator, ok := v.(scheme.Validator); ok {
		err := validator.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func decodeYaml(r *http.Request, v interface{}) error {
	if err := yaml.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	/* check for validity of the model*/
	if validator, ok := v.(scheme.Validator); ok {
		err := validator.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
