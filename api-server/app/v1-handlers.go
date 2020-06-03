package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chauhanr/singlenetes/api-server/scheme"
)

func (s *Server) StatusV1() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		msg := fmt.Sprintf("%s is: up", path)
		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Printf("Error writing status message %s\n", err)
		}
	}
}

func (s *Server) podHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Printf("URL: %s, pod configurations submitted to be saved.\n", path)
		data := scheme.ErrorMessage{}
		if r.Method == http.MethodPost {
			pod := scheme.PodV1{}
			err := DecodeYaml(r, &pod)
			if err != nil {
				msg := fmt.Sprintf("Error parsing PodV1 configuration %s\n", err)
				data.ParsingError(msg)
				respond(w, r, http.StatusInternalServerError, &data)
			}
			// add the pod definition to etcd cluster.
			podUid := guid()
			pod.Metadata.Uid = podUid
			namespace := "default"

		} else {
			data.MethodNotSupport(http.StatusMethodNotAllowed, r.Method)
			respond(w, r, http.StatusMethodNotAllowed, &data)
		}
	}
}
