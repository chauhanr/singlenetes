package app

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) StatusV1alpha() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		message := fmt.Sprint("%s status is: up", path)
		_, err := w.Write([]byte(message))
		if err != nil {
			log.Printf("Error writing status for /api/v2 %s\n", err)
		}
	}

}
