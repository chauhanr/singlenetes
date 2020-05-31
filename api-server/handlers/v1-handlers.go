package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func StatusV1(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	msg := fmt.Sprintf("%s is: up", path)
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Printf("Error writing status message %s\n", err)
	}
}
