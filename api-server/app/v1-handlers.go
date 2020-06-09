package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chauhanr/singlenetes/api-server/scheme"
	"github.com/gorilla/mux"
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

func (s *Server) podCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		vars := mux.Vars(r)
		log.Printf("URL: %s, pod configurations submitted to be saved.\n", path)
		data := scheme.ErrorMessage{}
		if r.Method == http.MethodPost {
			pod := scheme.PodV1{}
			err := DecodeYaml(r, &pod)
			if err != nil {
				msg := fmt.Sprintf("Error parsing PodV1 configuration %s\n", err)
				data.ParsingError(msg)
				respond(w, r, http.StatusInternalServerError, &data)
				return
			}
			// add the pod definition to etcd cluster.
			podUid := guid()
			pod.Metadata.Uid = podUid
			err = s.cli.AddPod(pod)
			pod.Metadata.Namespace = vars[NAMESPACE_PNAME]
			if err != nil {
				msg := fmt.Sprintf("Error saving PodV1 configuration %s\n", err)
				data.InternalServerError(http.StatusInternalServerError, msg)
				respond(w, r, http.StatusInternalServerError, &data)
				return
			}
			respond(w, r, http.StatusOK, nil)
		} else {
			data.MethodNotSupport(http.StatusMethodNotAllowed, r.Method)
			respond(w, r, http.StatusMethodNotAllowed, &data)
		}
	}
}

func (s *Server) registerEventSubscribers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Printf("URL: %s, Event Subscription Request\n", path)
		data := scheme.ErrorMessage{}
		if r.Method == http.MethodPost {
			subs := scheme.EventSubscriber{}
			err := DecodeYaml(r, &subs)
			if err != nil {
				msg := fmt.Sprintf("Error parsing Eventsubscriber %s\n", err)
				data.ParsingError(msg)
				respond(w, r, http.StatusInternalServerError, &data)
				return
			}
			err = s.cli.AddSubscriber(subs)
			if err != nil {
				msg := fmt.Sprintf("Error saving EventSubscriber config for %v error: %s\n", subs, err)
				data.InternalServerError(http.StatusInternalServerError, msg)
				respond(w, r, http.StatusInternalServerError, &data)
				return
			}
			respond(w, r, http.StatusOK, nil)
		} else {
			data.MethodNotSupport(http.StatusMethodNotAllowed, r.Method)
			respond(w, r, http.StatusMethodNotAllowed, &data)
		}

	}
}

func (s *Server) PodSubscribers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Printf("URL: %s, Get Subscribers List\n", path)
		vars := mux.Vars(r)
		data := scheme.ErrorMessage{}
		if r.Method == http.MethodGet {
			cmpType := vars[COMPONENT_TYPE]
			c, err := scheme.GetComponentType(cmpType)
			if err != nil {
				msg := fmt.Sprintf("Undefined Subscriber Type: %s\n", err)
				data.InternalServerError(http.StatusInternalServerError, msg)
				respond(w, r, http.StatusInternalServerError, &data)
				return
			}
			subscribers, err := s.cli.GetPodSubscribers(c)
			if err != nil {
				msg := fmt.Sprintf("Error retriving Pod Subscribers %s\n", err)
				data.InternalServerError(http.StatusInternalServerError, msg)
				respond(w, r, http.StatusInternalServerError, &data)
				return
			}
			respond(w, r, http.StatusOK, &subscribers)
			return
		} else {
			data.MethodNotSupport(http.StatusMethodNotAllowed, r.Method)
			respond(w, r, http.StatusMethodNotAllowed, &data)
		}
	}
}
